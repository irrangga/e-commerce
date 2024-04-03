package order

import (
	"time"
)

type Order struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ProductOrders []ProductOrder `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ProductOrder struct {
	ID        uint
	OrderID   uint
	ProductID uint
	Quantity  uint
}
