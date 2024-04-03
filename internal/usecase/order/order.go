package order

import (
	"context"
	"e-commerce/internal/entity"
)

func (uc Usecase) CreateOrder(ctx context.Context, input entity.CreateOrder) (entity.Order, error) {
	order, err := uc.repo.CreateOrder(ctx, input)
	if err != nil {
		return entity.Order{}, err
	}

	for index := range input.ProductOrders {
		input.ProductOrders[index].OrderID = order.ID
	}

	productOrders, err := uc.repo.CreateProductOrders(ctx, input)
	if err != nil {
		return entity.Order{}, err
	}
	order.ProductOrders = productOrders

	return order, nil
}

func (uc Usecase) DeleteOrder(ctx context.Context, id string) error {
	_, err := uc.repo.GetOrder(ctx, id)
	if err != nil {
		return err
	}

	err = uc.repo.DeleteProductOrderByOrderId(ctx, id)
	if err != nil {
		return err
	}

	err = uc.repo.DeleteOrder(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
