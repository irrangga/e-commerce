package entity

import "time"

type Order struct {
	ID            uint           `json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	ProductOrders []ProductOrder `json:"product_orders"`
}

type ProductOrder struct {
	OrderID   uint `json:"order_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type CreateOrder struct {
	ProductOrders []ProductOrder `json:"product_orders"`
}
