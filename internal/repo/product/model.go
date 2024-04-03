package product

import (
	"e-commerce/internal/repo/stock"
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Price     decimal.Decimal
	Stocks    []stock.Stock `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
