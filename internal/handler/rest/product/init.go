package product

import (
	"context"
	"e-commerce/internal/entity"
)

type UcInterface interface {
	GetListProducts(ctx context.Context) ([]entity.Product, error)
	GetProduct(ctx context.Context, id string) (entity.Product, error)
	CreateProduct(ctx context.Context, input entity.CreateProduct) (entity.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}

type Handler struct {
	uc UcInterface
}

func New(uc UcInterface) Handler {
	return Handler{
		uc: uc,
	}
}
