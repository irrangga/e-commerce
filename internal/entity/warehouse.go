package entity

import "time"

type Warehouse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	City      string    `json:"city"`
	Status    string    `json:"status"`
	Stocks    []Stock   `json:"stocks"`
}

type CreateWarehouse struct {
	City   string `json:"city"`
	Status string `json:"status"`
}

type UpdateWarehouse struct {
	ID     uint   `json:"id"`
	Status string `json:"status"`
}
