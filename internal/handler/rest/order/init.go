package order

import (
	"context"
	"e-commerce/internal/entity"
)

type UcInterface interface {
	CreateOrder(ctx context.Context, input entity.CreateOrder) (entity.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type Handler struct {
	uc UcInterface
}

func New(uc UcInterface) Handler {
	return Handler{
		uc: uc,
	}
}
