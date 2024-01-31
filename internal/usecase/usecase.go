package usecase

import (
	"context"

	"github.com/v-351/storage-api/internal/entity"
	"github.com/v-351/storage-api/internal/repository"
)

type Usecase struct {
	storage repository.Repo
}

func New(s repository.Repo) Usecase {
	return Usecase{
		storage: s,
	}
}

func (u Usecase) ReserveProduct(ctx context.Context, reserveOrder []entity.ReserveOrder) error {
	return u.storage.ReserveProduct(ctx, reserveOrder)
}

func (u Usecase) ReleaseReserveProcuct(ctx context.Context, reserveOrder []entity.ReserveOrder) error {
	return u.storage.ReleaseReserveProduct(ctx, reserveOrder)
}

func (u Usecase) GetWarehouseTotal(ctx context.Context, warehouseID uint) ([]entity.Placement, error) {
	return u.storage.GetWarehouseTotal(ctx, warehouseID)
}
