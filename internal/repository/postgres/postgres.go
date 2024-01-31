package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/v-351/storage-api/internal/config"
	"github.com/v-351/storage-api/internal/entity"
	"github.com/v-351/storage-api/internal/repository"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func New(config config.Config) Postgres {

	pool, err := pgxpool.New(context.Background(), config.PGConnString())
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return Postgres{pool: pool}
}

func (p Postgres) Close() {
	p.pool.Close()
}

func (db Postgres) GetWarehouseID(ctx context.Context, warehouse_name string) (uint, error) {
	const query = `
		SELECT id 
		FROM backend.warehouse
		WHERE name = $1`

	var ID uint
	err := db.pool.QueryRow(ctx, query, warehouse_name).Scan(&ID)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, fmt.Errorf("GetWarehouseID: %w", repository.ErrInvalidQueryArguments)
	}

	return ID, err
}

func (p Postgres) GetWarehouseTotal(ctx context.Context, warehouseID uint) ([]entity.Placement, error) {
	const query = `
		SELECT *
		FROM backend.placement 
		WHERE warehouse_id IN ($1)`

	rows, _ := p.pool.Query(ctx, query, warehouseID)
	defer rows.Close()

	placements, err := pgx.CollectRows[entity.Placement](rows, pgx.RowToStructByName[entity.Placement])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("GetWarehouseTotal: %w", repository.ErrInvalidQueryArguments)
	}

	return placements, nil
}

func (p Postgres) ReserveProduct(ctx context.Context, reserveOrder []entity.ReserveOrder) error {
	const query = `
		UPDATE backend.Placement 
		SET reserved = reserved + $1 
		WHERE warehouse_id = $2 AND product_id = $3`

	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("ReserveProduct: %w", err)
	}

	for _, o := range reserveOrder {
		_, err := tx.Exec(ctx, query, o.Reserved, o.WarehouseID, o.ProductID)
		if err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("ReserveProduct: %w", repository.ErrInvalidQueryArguments)
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("ReserveProduct: %w", err)
	}
	return nil
}

func (p Postgres) ReleaseReserveProduct(ctx context.Context, reserveOrder []entity.ReserveOrder) error {
	const query = `
		UPDATE backend.Placement 
		SET reserved = reserved - $1 
		WHERE warehouse_id = $2 AND product_id = $3`

	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("ReleaseReserveProduct: %w", err)
	}

	for _, o := range reserveOrder {
		_, err := tx.Exec(ctx, query, o.Reserved, o.WarehouseID, o.ProductID)
		if err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("ReleaseReserveProduct: %w", repository.ErrInvalidQueryArguments)
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("ReleaseReserveProduct: %w", err)
	}
	return nil
}
