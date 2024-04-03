package product

import (
	"context"
	"e-commerce/internal/entity"
)

func (r Repository) GetListProducts(ctx context.Context) ([]entity.Product, error) {
	var products []Product
	var productsEntity []entity.Product

	err := r.db.WithContext(ctx).Preload("Stocks").Find(&products).Error
	if err != nil {
		return []entity.Product{}, err
	}

	for _, product := range products {
		var stocksEntity []entity.Stock
		for _, stock := range product.Stocks {
			stocksEntity = append(stocksEntity, entity.Stock{
				ID:          stock.ID,
				CreatedAt:   stock.CreatedAt,
				UpdatedAt:   stock.UpdatedAt,
				ProductID:   stock.ProductID,
				WarehouseID: stock.WarehouseID,
				Quantity:    stock.Quantity,
			})
		}

		productsEntity = append(productsEntity, entity.Product{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
			Name:      product.Name,
			Price:     product.Price,
			Stocks:    stocksEntity,
		})
	}
	return productsEntity, nil
}

func (r Repository) GetProduct(ctx context.Context, id string) (entity.Product, error) {
	var product Product

	err := r.db.WithContext(ctx).Preload("Stocks").First(&product, id).Error
	if err != nil {
		return entity.Product{}, err
	}

	var stocksEntity []entity.Stock
	for _, stock := range product.Stocks {
		stocksEntity = append(stocksEntity, entity.Stock{
			ID:          stock.ID,
			CreatedAt:   stock.CreatedAt,
			UpdatedAt:   stock.UpdatedAt,
			ProductID:   stock.ProductID,
			WarehouseID: stock.WarehouseID,
			Quantity:    stock.Quantity,
		})
	}

	return entity.Product{
		ID:        product.ID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Name:      product.Name,
		Price:     product.Price,
		Stocks:    stocksEntity,
	}, nil
}

func (r Repository) CreateProduct(ctx context.Context, input entity.CreateProduct) (entity.Product, error) {
	product := Product{
		Name:  input.Name,
		Price: input.Price,
	}

	err := r.db.WithContext(ctx).Create(&product).Error
	if err != nil {
		return entity.Product{}, err
	}

	return entity.Product{
		ID:        product.ID,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Name:      product.Name,
		Price:     product.Price,
	}, nil
}

func (r Repository) DeleteProduct(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
