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

func (r Repository) DeleteStockByProductId(ctx context.Context, productId string) error {
	err := r.db.WithContext(ctx).Where("product_id = ?", productId).Delete(&Stock{}).Error
	if err != nil {
		return err
	}
	return nil
}
