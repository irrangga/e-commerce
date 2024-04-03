package product

import (
	"context"
	"e-commerce/internal/entity"
	"e-commerce/internal/usecase/stock"
)

type RepoInterface interface {
	GetListProducts(ctx context.Context) ([]entity.Product, error)
	GetProduct(ctx context.Context, id string) (entity.Product, error)
	CreateProduct(ctx context.Context, input entity.CreateProduct) (entity.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

type Usecase struct {
	productRepo RepoInterface
	stockRepo   stock.RepoInterface
}

func New(productRepo RepoInterface, stockRepo stock.RepoInterface) Usecase {
	return Usecase{
		productRepo: productRepo,
		stockRepo:   stockRepo,
	}
}
