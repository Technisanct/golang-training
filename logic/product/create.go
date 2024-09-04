package product

import (
	"context"
	"fmt"
)

func (p productImpl) Create(ctx context.Context, request *CreateProductRequest) error {
	
	fmt.Println("product created")

	return nil
}
