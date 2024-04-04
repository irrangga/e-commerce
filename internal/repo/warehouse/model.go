package warehouse

import (
	"e-commerce/internal/repo/stock"
	"time"
)

type Warehouse struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	City      string
	Status    string
	Stocks    []stock.Stock
}
