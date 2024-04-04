package order

import (
	"context"
	mock_order "e-commerce/gen/mocken/usecase/order"
	mock_stock "e-commerce/gen/mocken/usecase/stock"
	"e-commerce/internal/entity"
	"e-commerce/internal/usecase/stock"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestUsecase_CheckoutOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOrderRepoInterface := mock_order.NewMockRepoInterface(ctrl)
	mockStockRepoInterface := mock_stock.NewMockRepoInterface(ctrl)

	order := entity.Order{
		ID: 1,
		ProductOrders: []entity.ProductOrder{
			{
				OrderID:   1,
				ProductID: 1,
				Quantity:  20,
			},
		},
	}

	stocks := []entity.Stock{
		{
			ID:          1,
			ProductID:   1,
			WarehouseID: 1,
			Quantity:    25,
		},
		{
			ID:          2,
			ProductID:   2,
			WarehouseID: 1,
			Quantity:    10,
		},
	}

	type fields struct {
		orderRepo RepoInterface
		stockRepo stock.RepoInterface
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    entity.Order
		wantErr error
	}{
		// TEST CASES
		{
			name: "CheckoutOrder - Success",
			fields: fields{
				orderRepo: mockOrderRepoInterface,
				stockRepo: mockStockRepoInterface,
			},
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			mock: func() {
				mockOrderRepoInterface.EXPECT().GetOrder(gomock.Any(), gomock.Any()).
					Return(order, nil)

				mockStockRepoInterface.EXPECT().GetStocksByProductId(gomock.Any(), gomock.Any()).
					Return(stocks, nil)

				mockStockRepoInterface.EXPECT().UpdateStock(gomock.Any(), gomock.Any()).
					Return(entity.Stock{}, nil)
			},
			want:    order,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := Usecase{
				orderRepo: tt.fields.orderRepo,
				stockRepo: tt.fields.stockRepo,
			}
			tt.mock()
			got, err := uc.CheckoutOrder(tt.args.ctx, tt.args.id)
			if err != tt.wantErr {
				t.Errorf("Usecase.CheckoutOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.CheckoutOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
