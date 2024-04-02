package user

import (
	"context"
	"e-commerce/internal/entity"
)

type RepoInterface interface {
	GetUser(ctx context.Context, id string) (entity.User, error)
	CreateUser(ctx context.Context, input entity.CreateUser) (entity.User, error)
	UpdateUser(ctx context.Context, input entity.UpdateUser) (entity.User, error)
	DeleteUser(ctx context.Context, id string) (entity.User, error)
}

type Usecase struct {
	repo RepoInterface
}

func New(repo RepoInterface) Usecase {
	return Usecase{
		repo: repo,
	}
}
