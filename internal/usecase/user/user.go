package user

import (
	"context"
	"e-commerce/internal/entity"
	"strconv"
)

func (uc Usecase) GetUser(ctx context.Context, id string) (entity.User, error) {
	user, err := uc.repo.GetUser(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (uc Usecase) CreateUser(ctx context.Context, input entity.CreateUser) (entity.User, error) {
	user, err := uc.repo.CreateUser(ctx, input)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (uc Usecase) UpdateUser(ctx context.Context, input entity.UpdateUser) (entity.User, error) {
	_, err := uc.repo.GetUser(ctx, strconv.FormatUint(uint64(input.ID), 10))
	if err != nil {
		return entity.User{}, err
	}

	user, err := uc.repo.UpdateUser(ctx, input)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (uc Usecase) DeleteUser(ctx context.Context, id string) (entity.User, error) {
	_, err := uc.repo.GetUser(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	user, err := uc.repo.DeleteUser(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
