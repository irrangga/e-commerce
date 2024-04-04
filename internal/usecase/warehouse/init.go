package warehouse

import (
	"context"
	"e-commerce/internal/entity"
)

type RepoInterface interface {
	GetWarehouse(ctx context.Context, id string) (entity.Warehouse, error)
	CreateWarehouse(ctx context.Context, input entity.CreateWarehouse) (entity.Warehouse, error)
	UpdateWarehouse(ctx context.Context, input entity.UpdateWarehouse) (entity.Warehouse, error)
	DeleteWarehouse(ctx context.Context, id string) error
}

type Usecase struct {
	repo RepoInterface
}

func New(repo RepoInterface) Usecase {
	return Usecase{
		repo: repo,
	}
}
