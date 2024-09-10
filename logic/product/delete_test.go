package product

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	productMocks "golang-training/repository/product/mocks"
	"testing"
)

func Test_productImpl_Delete(t *testing.T) {
	type fields struct {
		product *productMocks.Product
	}
	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "happy case",
			fields: fields{product: mockDeleteProductRepo(nil)},
			args: args{
				ctx:  context.Background(),
				uuid: "uuid",
			},
			wantErr: nil,
		},

		{
			name:   "test_case fail",
			fields: fields{product: mockDeleteProductRepo(errors.New("product delete failed"))},
			args: args{
				ctx:  context.Background(),
				uuid: "uuid",
			},
			wantErr: errors.New("product delete failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := productImpl{
				product: tt.fields.product,
			}
			err := p.Delete(tt.args.ctx, tt.args.uuid)
			assert.Equal(t, tt.wantErr, err)
			tt.fields.product.AssertExpectations(t)
		})
	}
}

func mockDeleteProductRepo(createErr error) *productMocks.Product {
	client := &productMocks.Product{}
	client.On("DeleteOne", mock.Anything, mock.Anything).Return(createErr)

	return client
}
