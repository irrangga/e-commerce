package user

import (
	"context"
	"e-commerce/internal/entity"
)

type UcInterface interface {
	GetUser(ctx context.Context, id string) (entity.User, error)
	CreateUser(ctx context.Context, input entity.CreateUser) (entity.User, error)
	UpdateUser(ctx context.Context, input entity.UpdateUser) (entity.User, error)
	DeleteUser(ctx context.Context, id string) (entity.User, error)
}

type Handler struct {
	uc UcInterface
}

func New(uc UcInterface) Handler {
	return Handler{
		uc: uc,
	}
}
