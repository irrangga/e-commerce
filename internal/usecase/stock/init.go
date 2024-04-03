package stock

import (
	"context"
	"e-commerce/internal/entity"
)

type RepoInterface interface {
	CreateStock(ctx context.Context, input entity.CreateStock) (entity.Stock, error)
	DeleteStockByProductId(ctx context.Context, productId string) error
}

type Usecase struct {
	repo RepoInterface
}

func New(repo RepoInterface) Usecase {
	return Usecase{
		repo: repo,
	}
}
