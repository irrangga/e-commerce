package entity

import (
	"time"
)

type Product struct {
	ID                uint      `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Name              string    `json:"name"`
	StockAvailability uint      `json:"stock_availability"`
	Stocks            []Stock   `json:"stocks"`
}

type CreateProduct struct {
	Name   string  `json:"name"`
	Stocks []Stock `json:"stocks" binding:"dive"`
}
