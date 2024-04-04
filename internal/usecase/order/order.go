package order

import (
	"context"
	"e-commerce/internal/entity"
	"errors"
	"math"
	"strconv"
)

func (uc Usecase) CreateOrder(ctx context.Context, input entity.CreateOrder) (entity.Order, error) {
	order, err := uc.orderRepo.CreateOrder(ctx, input)
	if err != nil {
		return entity.Order{}, err
	}

	for index := range input.ProductOrders {
		input.ProductOrders[index].OrderID = order.ID
	}

	productOrders, err := uc.orderRepo.CreateProductOrders(ctx, input)
	if err != nil {
		return entity.Order{}, err
	}
	order.ProductOrders = productOrders

	return order, nil
}

func (uc Usecase) DeleteOrder(ctx context.Context, id string) error {
	_, err := uc.orderRepo.GetOrder(ctx, id)
	if err != nil {
		return err
	}

	err = uc.orderRepo.DeleteProductOrderByOrderId(ctx, id)
	if err != nil {
		return err
	}

	err = uc.orderRepo.DeleteOrder(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (uc Usecase) CheckoutOrder(ctx context.Context, id string) (entity.Order, error) {
	order, err := uc.orderRepo.GetOrder(ctx, id)
	if err != nil {
		return entity.Order{}, err
	}

	for _, productOrder := range order.ProductOrders {
		productId := strconv.FormatUint(uint64(productOrder.ProductID), 10)

		stocks, err := uc.stockRepo.GetStocksByProductId(ctx, productId)
		if err != nil {
			return entity.Order{}, err
		}

		stockAvailability := uint(0)
		for _, stock := range stocks {
			stockAvailability += stock.Quantity
		}

		if productOrder.Quantity > stockAvailability {
			return entity.Order{}, errors.New("product(s) out of stock")
		}

		quantity := productOrder.Quantity
		for _, stock := range stocks {
			remainderQuantity := stock.Quantity - quantity

			if remainderQuantity <= 0 {
				_, err := uc.stockRepo.UpdateStock(ctx, entity.UpdateStock{
					ID:       stock.ID,
					Quantity: 0,
				})
				if err != nil {
					return entity.Order{}, err
				}
				quantity = uint(math.Abs(float64(remainderQuantity)))

			} else {
				_, err := uc.stockRepo.UpdateStock(ctx, entity.UpdateStock{
					ID:       stock.ID,
					Quantity: remainderQuantity,
				})
				if err != nil {
					return entity.Order{}, err
				}
				break
			}
		}
	}

	return order, nil
}
