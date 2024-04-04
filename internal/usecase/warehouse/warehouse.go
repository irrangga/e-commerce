package warehouse

import (
	"context"
	"e-commerce/internal/entity"
	"strconv"
)

func (uc Usecase) GetWarehouse(ctx context.Context, id string) (entity.Warehouse, error) {
	warehouse, err := uc.repo.GetWarehouse(ctx, id)
	if err != nil {
		return entity.Warehouse{}, err
	}
	return warehouse, nil
}

func (uc Usecase) CreateWarehouse(ctx context.Context, input entity.CreateWarehouse) (entity.Warehouse, error) {
	warehouse, err := uc.repo.CreateWarehouse(ctx, input)
	if err != nil {
		return entity.Warehouse{}, err
	}
	return warehouse, nil
}

func (uc Usecase) DeleteWarehouse(ctx context.Context, id string) error {
	_, err := uc.repo.GetWarehouse(ctx, id)
	if err != nil {
		return err
	}

	err = uc.repo.DeleteWarehouse(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (uc Usecase) UpdateStatusWarehouse(ctx context.Context, input entity.UpdateWarehouse) (entity.Warehouse, error) {
	warehouse, err := uc.repo.GetWarehouse(ctx, strconv.FormatUint(uint64(input.ID), 10))
	if err != nil {
		return entity.Warehouse{}, err
	}

	updatedWarehouse, err := uc.repo.UpdateWarehouse(ctx, input)
	if err != nil {
		return entity.Warehouse{}, err
	}
	warehouse.Status = updatedWarehouse.Status

	return warehouse, nil
}
