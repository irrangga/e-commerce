package product

import (
	"context"
	"e-commerce/internal/entity"
)

func (uc Usecase) GetListProducts(ctx context.Context) ([]entity.Product, error) {
	products, err := uc.productRepo.GetListProducts(ctx)
	if err != nil {
		return []entity.Product{}, err
	}

	for index, product := range products {
		stockAvailability := uint(0)
		for _, stock := range product.Stocks {
			stockAvailability += stock.Quantity
		}
		products[index].StockAvailability = stockAvailability
	}

	return products, nil
}

func (uc Usecase) GetProduct(ctx context.Context, id string) (entity.Product, error) {
	product, err := uc.productRepo.GetProduct(ctx, id)
	if err != nil {
		return entity.Product{}, err
	}

	stockAvailability := uint(0)
	for _, stock := range product.Stocks {
		stockAvailability += stock.Quantity
	}
	product.StockAvailability = stockAvailability

	return product, nil
}

func (uc Usecase) CreateProduct(ctx context.Context, input entity.CreateProduct) (entity.Product, error) {
	product, err := uc.productRepo.CreateProduct(ctx, input)
	if err != nil {
		return entity.Product{}, err
	}

	var stocksEntity []entity.Stock
	for _, stock := range input.Stocks {
		stock, err := uc.stockRepo.CreateStock(ctx, entity.CreateStock{
			ProductID:   product.ID,
			WarehouseID: stock.WarehouseID,
			Quantity:    stock.Quantity,
		})
		if err != nil {
			return entity.Product{}, err
		}

		stocksEntity = append(stocksEntity, stock)
	}
	product.Stocks = stocksEntity

	return product, nil
}

func (uc Usecase) DeleteProduct(ctx context.Context, id string) error {
	_, err := uc.productRepo.GetProduct(ctx, id)
	if err != nil {
		return err
	}

	err = uc.stockRepo.DeleteStockByProductId(ctx, id)
	if err != nil {
		return err
	}

	err = uc.productRepo.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
