package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/v-351/storage-api/internal/entity"
	"github.com/v-351/storage-api/internal/repository"
)

// @Summary      Содержимое склада
// @Description  Получить содержимое склада по идентификатору
// @Produce      json
// @Param        warehouse_id   path      uint  true  "ID Склада"
// @Success      200  {object}  warehouseTotalResponse
// @Failure      400
// @Failure      500
// @Router       /warehouse/{warehouse_id}/total [get]
func (s *Server) getWarehouseTotal(c echo.Context) error {
	form := new(warehouseForm)
	if err := c.Bind(form); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(form); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	result, err := s.usecase.GetWarehouseTotal(c.Request().Context(), form.WarehouseID)
	if err != nil {
		s.ll.Error(err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	responseData := make([]placement, 0, len(result))
	for _, r := range result {
		responseData = append(responseData,
			placement{
				ProductID: r.ProductID,
				Total:     r.Total,
				Reserved:  r.Reserved,
			},
		)
	}

	return c.JSON(http.StatusOK, warehouseTotalResponse{Placement: responseData})
}

// @Summary      Зарезервирование товара
// @Description  Зарезервировать товар на складе
// @Accept       json
// @Param        reserveOrder   body  reserveOrderForm true  "Список товаров для резервирования"
// @Success      201
// @Failure      400
// @Failure      500
// @Router       /reserve [post]
func (s *Server) reserveProduct(c echo.Context) error {
	form := new(reserveOrderForm)
	if err := c.Bind(form); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(form); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	orders := make([]entity.ReserveOrder, 0, len(form.ReserveOrder))
	for _, ro := range form.ReserveOrder {
		orders = append(orders, entity.ReserveOrder{
			WarehouseID: ro.WarehouseID,
			ProductID:   ro.ProductID,
			Reserved:    ro.ReserveCount,
		})
	}

	err := s.usecase.ReserveProduct(c.Request().Context(), orders)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidQueryArguments) {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}

// @Summary      Снятие резерва товара
// @Description  Снятие резервирования с товар на складе
// @Accept       json
// @Param        reserveOrder   body  reserveOrderForm true  "Список товаров для разрезервирования"
// @Success      201
// @Failure      400
// @Failure      500
// @Router       /release [post]
func (s *Server) releaseReserveProcuct(c echo.Context) error {
	form := new(reserveOrderForm)
	if err := c.Bind(form); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	if err := c.Validate(form); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	orders := make([]entity.ReserveOrder, 0, len(form.ReserveOrder))
	for _, ro := range form.ReserveOrder {
		orders = append(orders, entity.ReserveOrder{
			WarehouseID: ro.WarehouseID,
			ProductID:   ro.ProductID,
			Reserved:    ro.ReserveCount,
		})
	}

	err := s.usecase.ReleaseReserveProcuct(c.Request().Context(), orders)
	if err != nil {
		if errors.Is(err, repository.ErrInvalidQueryArguments) {
			return c.NoContent(http.StatusBadRequest)
		}
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}
