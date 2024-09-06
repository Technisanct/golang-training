package product

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	productMocks "golang-training/repository/product/mocks"
	"testing"
)

func Test_productImpl_Create(t *testing.T) {
	type fields struct {
		product *productMocks.Product
	}
	type args struct {
		ctx     context.Context
		request *CreateProductRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "test_case_pass",
			fields: fields{product: mockCreateProductRepo(true, nil)},
			args: args{
				ctx: context.Background(),
				request: &CreateProductRequest{
					Name:  "p1",
					Price: 1200,
				},
			},
			wantErr: nil,
		},
		{
			name:   "test_case_fail",
			fields: fields{product: mockCreateProductRepo(true, errors.New("Failed"))},
			args: args{
				ctx: context.Background(),
				request: &CreateProductRequest{
					Name:  "p2",
					Price: 200,
				},
			},
			wantErr: errors.New("Failed"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := productImpl{
				product: tt.fields.product,
			}
			err := p.Create(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err)

			tt.fields.product.AssertExpectations(t)
		})
	}
}

func mockCreateProductRepo(enableFlag bool, createErr error) *productMocks.Product {
	client := &productMocks.Product{}
	if enableFlag {
		client.On("Create", mock.Anything, mock.Anything).Return(createErr)
	}

	return client
}
