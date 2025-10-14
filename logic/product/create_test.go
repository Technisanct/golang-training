package product

import (
	"context"
	"errors"
	"golang-training/logic/product/contract"
	"golang-training/repository/product/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	createProductRequest = contract.CreateProductRequest{
		Name:            "test-1",
		Price:           100.00,
		DiscountedPrice: 10.00,
	}
	invalidCreateProductRequest = contract.CreateProductRequest{
		Name:            "test-1",
		DiscountedPrice: 10.00,
	}
)

func TestProductService(t *testing.T) {
	type fields struct {
		product *mocks.Product
	}
	type Args struct {
		data *contract.CreateProductRequest
		ctx  context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    Args
		wantErr error
	}{
		{
			name:   "happy path",
			fields: fields{product: mockCreateProductRepo(true, nil)},
			args: Args{
				data: &createProductRequest,
				ctx:  context.Background(),
			},
			wantErr: nil,
		},
		{
			name:   "repo error",
			fields: fields{product: mockCreateProductRepo(true, errors.New("failed"))},
			args: Args{
				data: &invalidCreateProductRequest,
				ctx:  context.Background(),
			},
			wantErr: errors.New("failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prod := productImpl{
				product: tt.fields.product,
			}
			err := prod.Create(tt.args.ctx, tt.args.data)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func mockCreateProductRepo(enableFlag bool, createErr error) *mocks.Product {
	client := &mocks.Product{}
	if enableFlag {
		client.
			On("Create", mock.Anything, mock.Anything).
			Return(createErr)
	}

	return client
}
