package warehouse

import (
	"context"
	"e-commerce/internal/entity"
)

func (r Repository) GetWarehouse(ctx context.Context, id string) (entity.Warehouse, error) {
	var warehouse Warehouse

	err := r.db.WithContext(ctx).Preload("Stocks").First(&warehouse, id).Error
	if err != nil {
		return entity.Warehouse{}, err
	}

	var stocksEntity []entity.Stock
	for _, stock := range warehouse.Stocks {
		stocksEntity = append(stocksEntity, entity.Stock{
			ID:          stock.ID,
			CreatedAt:   stock.CreatedAt,
			UpdatedAt:   stock.UpdatedAt,
			ProductID:   stock.ProductID,
			WarehouseID: stock.WarehouseID,
			Quantity:    stock.Quantity,
		})
	}

	return entity.Warehouse{
		ID:        warehouse.ID,
		CreatedAt: warehouse.CreatedAt,
		UpdatedAt: warehouse.UpdatedAt,
		City:      warehouse.City,
		Stocks:    stocksEntity,
	}, nil
}

func (r Repository) CreateWarehouse(ctx context.Context, input entity.CreateWarehouse) (entity.Warehouse, error) {
	warehouse := Warehouse{
		City:   input.City,
		Status: input.Status,
	}

	err := r.db.WithContext(ctx).Create(&warehouse).Error
	if err != nil {
		return entity.Warehouse{}, err
	}
	return entity.Warehouse{
		ID:        warehouse.ID,
		CreatedAt: warehouse.CreatedAt,
		UpdatedAt: warehouse.UpdatedAt,
		City:      warehouse.City,
		Status:    warehouse.Status,
	}, nil
}

func (r Repository) UpdateWarehouse(ctx context.Context, input entity.UpdateWarehouse) (entity.Warehouse, error) {
	warehouse := Warehouse{
		ID:     input.ID,
		Status: input.Status,
	}

	err := r.db.WithContext(ctx).Model(&warehouse).Update("status", input.Status).Error
	if err != nil {
		return entity.Warehouse{}, err
	}

	return entity.Warehouse{
		ID:        warehouse.ID,
		CreatedAt: warehouse.CreatedAt,
		UpdatedAt: warehouse.UpdatedAt,
		Status:    warehouse.Status,
	}, nil
}

func (r Repository) DeleteWarehouse(ctx context.Context, id string) error {
	err := r.db.WithContext(ctx).Delete(&Warehouse{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
