package product

import (
	"context"
	"golang-training/repository/product"
	"testing"
)

func Test_productImpl_Create(t *testing.T) {
	type fields struct {
		product product.Product
	}
	type args struct {
		ctx     context.Context
		request *CreateProductRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := productImpl{
				product: tt.fields.product,
			}
			if err := p.Create(tt.args.ctx, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
