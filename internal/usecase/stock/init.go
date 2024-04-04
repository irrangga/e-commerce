package stock

import (
	"context"
	"e-commerce/internal/entity"
)

type RepoInterface interface {
	CreateStock(ctx context.Context, input entity.CreateStock) (entity.Stock, error)
	DeleteStocksByProductId(ctx context.Context, productId string) error
	UpdateStock(ctx context.Context, input entity.UpdateStock) (entity.Stock, error)
	GetStocksByProductId(ctx context.Context, productId string) ([]entity.Stock, error)
}

type Usecase struct {
	repo RepoInterface
}

func New(repo RepoInterface) Usecase {
	return Usecase{
		repo: repo,
	}
}
