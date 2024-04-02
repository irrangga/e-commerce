package user

import (
	"context"
	"e-commerce/internal/entity"
)

func (r Repository) GetUser(ctx context.Context, id string) (entity.User, error) {
	var user User

	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}, nil
}

func (r Repository) CreateUser(ctx context.Context, input entity.CreateUser) (entity.User, error) {
	user := User{
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}, nil
}

func (r Repository) UpdateUser(ctx context.Context, input entity.UpdateUser) (entity.User, error) {
	user := User{
		ID:          input.ID,
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}

	err := r.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{
		ID:          user.ID,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
	}, nil
}

func (r Repository) DeleteUser(ctx context.Context, id string) (entity.User, error) {
	var user User

	err := r.db.WithContext(ctx).Delete(&user, id).Error
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{}, nil
}
