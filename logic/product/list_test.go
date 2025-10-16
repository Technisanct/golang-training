package product

import (
	"context"
	"errors"
	"golang-training/repository/model"
	"golang-training/repository/product/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	returnListProductData = []model.Product{
		{
			Name:            "test-1",
			Price:           100.000,
			DiscountedPrice: 10.00,
		},
		{
			Name:            "test-1",
			Price:           100.000,
			DiscountedPrice: 10.00,
		},
	}
)

func TestListProducts(t *testing.T) {
	type fields struct {
		product *mocks.Product
	}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name             string
		args             args
		fields           fields
		wantListProducts []model.Product
		wantErr          error
	}{
		{
			name: "happy path",
			args: args{
				ctx: context.Background(),
			},
			fields: fields{
				product: mockListProductRepo(true, returnListProductData, nil),
			},
			wantErr:          nil,
			wantListProducts: returnListProductData,
		},
		{
			name: "error",
			args: args{
				ctx: context.Background(),
			},
			fields: fields{
				product: mockListProductRepo(true, nil, errors.New("failed to fetch")),
			},
			wantErr:          errors.New("failed to fetch"),
			wantListProducts: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := productImpl{
				repo: tt.fields.product,
			}
			products, err := p.List(tt.args.ctx)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.wantListProducts, products)
		})
	}
}

func mockListProductRepo(enableFlag bool, returnProductListData []model.Product, createErr error) *mocks.Product {
	client := &mocks.Product{}
	if enableFlag {
		client.On("List", mock.Anything).Return(
			returnProductListData,
			createErr,
		)
	}
	return client
}
