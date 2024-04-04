package stock

import (
	"context"
	"e-commerce/internal/entity"
)

func (r Repository) CreateStock(ctx context.Context, input entity.CreateStock) (entity.Stock, error) {
	stock := Stock{
		ProductID:   input.ProductID,
		WarehouseID: input.WarehouseID,
		Quantity:    input.Quantity,
	}

	err := r.db.WithContext(ctx).Create(&stock).Error
	if err != nil {
		return entity.Stock{}, err
	}

	return entity.Stock{
		ID:          stock.ID,
		CreatedAt:   stock.CreatedAt,
		UpdatedAt:   stock.UpdatedAt,
		ProductID:   stock.ProductID,
		WarehouseID: stock.WarehouseID,
		Quantity:    stock.Quantity,
	}, nil
}

func (r Repository) DeleteStocksByProductId(ctx context.Context, productId string) error {
	err := r.db.WithContext(ctx).Where("product_id = ?", productId).Delete(&Stock{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) UpdateStock(ctx context.Context, input entity.UpdateStock) (entity.Stock, error) {
	stock := Stock{
		ID:       input.ID,
		Quantity: input.Quantity,
	}

	err := r.db.WithContext(ctx).Model(&stock).Update("quantity", input.Quantity).Error
	if err != nil {
		return entity.Stock{}, err
	}

	return entity.Stock{
		ID:          stock.ID,
		CreatedAt:   stock.CreatedAt,
		UpdatedAt:   stock.UpdatedAt,
		ProductID:   stock.ProductID,
		WarehouseID: stock.WarehouseID,
		Quantity:    stock.Quantity,
	}, nil
}

func (r Repository) GetStocksByProductId(ctx context.Context, productId string) ([]entity.Stock, error) {
	var stocks []Stock
	var stocksEntity []entity.Stock

	err := r.db.WithContext(ctx).Where("product_id = ?", productId).Find(&stocks).Error
	if err != nil {
		return []entity.Stock{}, err
	}

	for _, stock := range stocks {
		stocksEntity = append(stocksEntity, entity.Stock{
			ID:          stock.ID,
			CreatedAt:   stock.CreatedAt,
			UpdatedAt:   stock.UpdatedAt,
			ProductID:   stock.ProductID,
			WarehouseID: stock.WarehouseID,
			Quantity:    stock.Quantity,
		})
	}

	return stocksEntity, nil
}
