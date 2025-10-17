package product

import (
	"context"
	"errors"
	"golang-training/logic/product/contract"
	"golang-training/repository/model"
	"golang-training/repository/product/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	fixedObjectId = primitive.NewObjectID()
	fixedTestTime = time.Date(2025, time.October, 10, 0, 0, 0, 0, time.UTC)

	returnListProductData = []model.Product{
		{
			ID:              &fixedObjectId,
			Name:            "test-1",
			Price:           100.000,
			DiscountedPrice: 10.00,
			CreatedAt:       fixedTestTime,
			UpdatedAt:       fixedTestTime,
		},
		{
			ID:              &fixedObjectId,
			Name:            "test-1",
			Price:           100.000,
			DiscountedPrice: 10.00,
			CreatedAt:       fixedTestTime,
			UpdatedAt:       fixedTestTime,
		},
	}

	expectedProductOutput = []contract.Product{
		{
			ID:              fixedObjectId.String(),
			Name:            "test-1",
			Price:           100.000,
			DiscountedPrice: 10.00,
			CreatedAt:       fixedTestTime,
			UpdatedAt:       fixedTestTime,
		},
		{
			ID:              fixedObjectId.String(),
			Name:            "test-1",
			Price:           100.000,
			DiscountedPrice: 10.00,
			CreatedAt:       fixedTestTime,
			UpdatedAt:       fixedTestTime,
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
		wantListProducts []contract.Product
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
			wantListProducts: expectedProductOutput,
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
