package stock

import "time"

type Stock struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ProductID   uint
	WarehouseID uint
	Quantity    uint
}
