package order

import (
	"context"
	"e-commerce/internal/entity"
	"e-commerce/internal/usecase/stock"
)

type RepoInterface interface {
	GetOrder(ctx context.Context, id string) (entity.Order, error)
	CreateOrder(ctx context.Context, input entity.CreateOrder) (entity.Order, error)
	CreateProductOrders(ctx context.Context, input entity.CreateOrder) ([]entity.ProductOrder, error)
	DeleteOrder(ctx context.Context, id string) error
	DeleteProductOrderByOrderId(ctx context.Context, orderId string) error
}

type Usecase struct {
	orderRepo RepoInterface
	stockRepo stock.RepoInterface
}

func New(orderRepo RepoInterface, stockRepo stock.RepoInterface) Usecase {
	return Usecase{
		orderRepo: orderRepo,
		stockRepo: stockRepo,
	}
}
