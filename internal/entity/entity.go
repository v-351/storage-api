package entity

type Warehouse struct {
	ID            uint   `db:"id"`
	Name          string `db:"name"`
	Accessibility bool   `db:"accessibility"`
}

type Product struct {
	Name string `db:"name"`
	Mass uint   `db:"mass"`
	ID   uint   `db:"id"`
}

type Placement struct {
	WarehouseID uint `db:"warehouse_id"`
	ProductID   uint `db:"product_id"`
	Total       uint `db:"total"`
	Reserved    uint `db:"reserved"`
}

type ReserveOrder struct {
	WarehouseID uint
	ProductID   uint
	Reserved    uint
}
