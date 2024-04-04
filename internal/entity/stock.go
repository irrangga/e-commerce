package entity

import (
	"time"
)

type Stock struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	ProductID   uint      `json:"product_id"`
	WarehouseID uint      `json:"warehouse_id"`
	Quantity    uint      `json:"quantity"`
}

type CreateStock struct {
	ProductID   uint `json:"product_id"`
	WarehouseID uint `json:"warehouse_id"`
	Quantity    uint `json:"quantity"`
}

type UpdateStock struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}
