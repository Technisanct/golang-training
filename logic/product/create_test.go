package product

import (
	"context"
	"errors"
	"golang-training/logic/product/contract"
	"golang-training/repository/product/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductService(t *testing.T) {

	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	type mockRepo struct {
		product *mocks.Product
	}
	type Args struct {
		data *contract.CreateProductRequest
	}

	payload := &contract.CreateProductRequest{
		Name:            "test-1",
		Price:           100.00,
		DiscountedPrice: 10.00,
	}
	tests := []struct {
		name    string
		mocked  mockRepo
		ctx     context.Context
		args    Args
		wantErr error
	}{
		{
			name:   "happy path",
			ctx:    c,
			mocked: mockRepo{product: mockCreateProductRepo(true, nil)},
			args: Args{
				data: payload,
			},
			wantErr: nil,
		},
		{
			name:   "repo error",
			ctx:    c,
			mocked: mockRepo{product: mockCreateProductRepo(true, errors.New("failed"))},
			args: Args{
				data: &contract.CreateProductRequest{
					Name:            "naruto-x",
					DiscountedPrice: 10.00,
				},
			},
			wantErr: errors.New("failed"),
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			prod := productImpl{
				product: tt.mocked.product,
			}
			err := prod.Create(tt.ctx, tt.args.data)
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
