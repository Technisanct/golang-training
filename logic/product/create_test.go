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
	type args struct {
		request *contract.CreateProductRequest
		ctx     context.Context
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "happy path",
			fields: fields{product: mockCreateProductRepo(true, nil)},
			args: args{
				request: &createProductRequest,
				ctx:     context.Background(),
			},
			wantErr: nil,
		},
		{
			name:   "repo error",
			fields: fields{product: mockCreateProductRepo(true, errors.New("failed"))},
			args: args{
				request: &invalidCreateProductRequest,
				ctx:     context.Background(),
			},
			wantErr: errors.New("failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := productImpl{
				product: tt.fields.product,
			}
			err := p.Create(tt.args.ctx, tt.args.request)
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
