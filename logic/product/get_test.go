package product

import (
	"context"
	"errors"
	"fmt"
	"golang-training/logic/product/contract"
	"golang-training/repository/model"
	"golang-training/repository/product/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	objectId                      = "68e8a5a61b2689c270bf176a"
	expectedProductOutputFromRepo = model.Product{
		ID:              &fixedObjectId,
		Name:            "test-1",
		Price:           100.000,
		DiscountedPrice: 10.00,
		CreatedAt:       fixedTestTime,
		UpdatedAt:       fixedTestTime,
	}

	expectedProductOutputFromLogic = mapSingleProductFromRepoToLogic(&expectedProductOutputFromRepo)
)

func TestGetProduct(t *testing.T) {
	type fields struct {
		product *mocks.Product
	}
	type args struct {
		ctx context.Context
	}

	tests := []struct {
		name           string
		args           args
		fields         fields
		wantErr        error
		expectedOutput *contract.Product
	}{
		{
			name: "happy path",
			args: args{
				ctx: context.Background(),
			},
			fields: fields{
				product: mockGetProductRepo(true, &expectedProductOutputFromRepo, nil),
			},
			expectedOutput: expectedProductOutputFromLogic,
			wantErr:        nil,
		},
		{
			name: "logic error",
			args: args{
				ctx: context.Background(),
			},
			fields: fields{
				product: mockGetProductRepo(true, nil, errors.New("failed to fetch product")),
			},
			wantErr:        errors.New("failed to fetch product"),
			expectedOutput: nil,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			p := productImpl{
				repo: tt.fields.product,
			}

			productOutput, err := p.Get(tt.args.ctx, fixedObjectId.Hex())
			fmt.Println("output", productOutput)
			assert.Equal(t, err, tt.wantErr)
			assert.Equal(t, tt.expectedOutput, productOutput)

		})

	}

}

func mockGetProductRepo(enableFlag bool, returnProductData *model.Product, createErr error) *mocks.Product {
	client := &mocks.Product{}

	if enableFlag {
		client.
			On("Get", mock.Anything, mock.Anything).
			Return(returnProductData, createErr)
	}

	return client
}
