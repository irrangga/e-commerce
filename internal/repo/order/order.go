package order

import (
	"context"
	"e-commerce/internal/entity"
)

func (r Repository) GetOrder(ctx context.Context, id string) (entity.Order, error) {
	var order Order

	err := r.db.WithContext(ctx).Preload("ProductOrders").First(&order, id).Error
	if err != nil {
		return entity.Order{}, err
	}

	var productOrdersEntity []entity.ProductOrder
	for _, productOrder := range order.ProductOrders {
		productOrdersEntity = append(productOrdersEntity, entity.ProductOrder{
			OrderID:   productOrder.OrderID,
			ProductID: productOrder.ProductID,
			Quantity:  productOrder.Quantity,
		})
	}

	return entity.Order{
		ID:            order.ID,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
		ProductOrders: productOrdersEntity,
	}, nil
}

func (r Repository) CreateOrder(ctx context.Context, input entity.CreateOrder) (entity.Order, error) {
	var order Order

	err := r.db.WithContext(ctx).Create(&order).Error
	if err != nil {
		return entity.Order{}, err
	}

	return entity.Order{
		ID:        order.ID,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}, nil
}

func (r Repository) CreateProductOrders(ctx context.Context, input entity.CreateOrder) ([]entity.ProductOrder, error) {
	var productOrders []ProductOrder

	for _, productOrder := range input.ProductOrders {
		productOrders = append(productOrders, ProductOrder{
			OrderID:   productOrder.OrderID,
			ProductID: productOrder.ProductID,
			Quantity:  productOrder.Quantity,
		})
	}

	err := r.db.WithContext(ctx).Create(&productOrders).Error
	if err != nil {
		return []entity.ProductOrder{}, err
	}

	var productOrdersEntity []entity.ProductOrder
	for _, productOrder := range productOrders {
		productOrdersEntity = append(productOrdersEntity, entity.ProductOrder{
			OrderID:   productOrder.OrderID,
			ProductID: productOrder.ProductID,
			Quantity:  productOrder.Quantity,
		})
	}

	return productOrdersEntity, nil
}

func (r Repository) DeleteOrder(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&Order{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) DeleteProductOrderByOrderId(ctx context.Context, orderId string) error {
	err := r.db.WithContext(ctx).Where("product_id = ?", orderId).Delete(&ProductOrder{}).Error
	if err != nil {
		return err
	}
	return nil
}
