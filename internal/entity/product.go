package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID                uint            `json:"id"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	Name              string          `json:"name"`
	Price             decimal.Decimal `json:"price"`
	StockAvailability uint            `json:"stock_availability"`
	Stocks            []Stock         `json:"stocks"`
}

type CreateProduct struct {
	Name   string          `json:"name"`
	Price  decimal.Decimal `json:"price"`
	Stocks []Stock         `json:"stocks" binding:"dive"`
}
