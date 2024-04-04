package warehouse

import (
	"context"
	"e-commerce/internal/entity"
)

type UcInterface interface {
	GetWarehouse(ctx context.Context, id string) (entity.Warehouse, error)
	CreateWarehouse(ctx context.Context, input entity.CreateWarehouse) (entity.Warehouse, error)
	DeleteWarehouse(ctx context.Context, id string) error
	UpdateStatusWarehouse(ctx context.Context, input entity.UpdateWarehouse) (entity.Warehouse, error)
}

type Handler struct {
	uc UcInterface
}

func New(uc UcInterface) Handler {
	return Handler{
		uc: uc,
	}
}
