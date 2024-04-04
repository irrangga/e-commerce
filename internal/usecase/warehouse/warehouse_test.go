package warehouse

import (
	"context"
	mock_warehouse "e-commerce/gen/mocken/usecase/warehouse"
	"e-commerce/internal/entity"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestUsecase_UpdateStatusWarehouse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWarehouseRepoInterface := mock_warehouse.NewMockRepoInterface(ctrl)

	warehouse := entity.Warehouse{
		ID:     1,
		City:   "Jakarta",
		Status: "ACTIVE",
	}

	updatedWarehouse := entity.Warehouse{
		ID:     1,
		Status: "INACTIVE",
	}

	type fields struct {
		repo RepoInterface
	}
	type args struct {
		ctx   context.Context
		input entity.UpdateWarehouse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Warehouse
		wantErr error
	}{
		// TEST CASES
		{
			name: "UpdateStatusWarehouse - Success",
			fields: fields{
				repo: mockWarehouseRepoInterface,
			},
			args: args{
				ctx: context.Background(),
				input: entity.UpdateWarehouse{
					ID:     1,
					Status: "INACTIVE",
				},
			},
			mock: func() {
				mockWarehouseRepoInterface.EXPECT().GetWarehouse(gomock.Any(), gomock.Any()).Return(warehouse, nil)
				mockWarehouseRepoInterface.EXPECT().UpdateWarehouse(gomock.Any(), gomock.Any()).Return(updatedWarehouse, nil)
			},
			want: entity.Warehouse{
				ID:     1,
				City:   "Jakarta",
				Status: "INACTIVE",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := Usecase{
				repo: tt.fields.repo,
			}
			tt.mock()
			got, err := uc.UpdateStatusWarehouse(tt.args.ctx, tt.args.input)
			if err != tt.wantErr {
				t.Errorf("Usecase.UpdateStatusWarehouse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.UpdateStatusWarehouse() = %v, want %v", got, tt.want)
			}
		})
	}
}
