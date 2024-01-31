package repository

import (
	"context"
	"errors"

	"github.com/v-351/storage-api/internal/entity"
)

type Repo interface {
	GetWarehouseTotal(context.Context, uint) ([]entity.Placement, error)
	ReserveProduct(context.Context, []entity.ReserveOrder) error
	ReleaseReserveProduct(context.Context, []entity.ReserveOrder) error
}

var ErrInvalidQueryArguments = errors.New("invalid query arguments")
