package http

type warehouseForm struct {
	WarehouseID uint `param:"warehouse_id" validate:"ne=0"`
}

type reserveOrder struct {
	WarehouseID  uint `json:"warehouse_id" validate:"ne=0"`
	ProductID    uint `json:"product_id" validate:"ne=0"`
	ReserveCount uint `json:"reserve_count" validate:"ne=0"`
}

type reserveOrderForm struct {
	ReserveOrder []reserveOrder `json:"reserve_order"`
}
