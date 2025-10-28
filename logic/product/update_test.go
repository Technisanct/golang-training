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
	productId          = "68e8a5a61b2689c270bf176a"
	ProductRequestData = contract.UpdateProductRequest{
		Name:            "test-1",
		Price:           200,
		DiscountedPrice: 10.00,
	}
)

func TestLogicUpdateProduct(t *testing.T) {
	type fields struct {
		product *mocks.Product
	}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            error
		productRequestData *contract.UpdateProductRequest
	}{
		{
			name: "happy path",
			fields: fields{
				product: mockUpdateProductRepo(true, nil),
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr:            nil,
			productRequestData: &ProductRequestData,
		},
		{
			name: "logic error",
			fields: fields{
				product: mockUpdateProductRepo(true, errors.New("product not found")),
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr:            errors.New("product not found"),
			productRequestData: &ProductRequestData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := productImpl{
				repo: tt.fields.product,
			}
			err := p.Update(tt.args.ctx, productId, tt.productRequestData)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func mockUpdateProductRepo(enableFlag bool, createErr error) *mocks.Product {
	client := &mocks.Product{}
	if enableFlag {
		client.On("Update", mock.Anything, mock.Anything, mock.Anything).
			Return(createErr)
	}

	return client
}
