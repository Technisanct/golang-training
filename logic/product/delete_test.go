package product

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang-training/repository/product"
	"testing"
)

func Test_productImpl_Delete(t *testing.T) {
	type fields struct {
		product product.Product
	}
	type args struct {
		ctx  context.Context
		uuid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := productImpl{
				product: tt.fields.product,
			}
			tt.wantErr(t, p.Delete(tt.args.ctx, tt.args.uuid), fmt.Sprintf("Delete(%v, %v)", tt.args.ctx, tt.args.uuid))
		})
	}
}