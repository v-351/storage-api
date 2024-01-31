package http

type placement struct {
	ProductID uint `json:"product_id"`
	Total     uint `json:"total"`
	Reserved  uint `json:"reserved"`
}

type warehouseTotalResponse struct {
	Placement []placement `json:"placement"`
}
