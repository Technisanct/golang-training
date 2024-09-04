package product

import "context"

type Products interface {
	Create(ctx context.Context, request *CreateProductRequest) error
}

type productImpl struct{}

func New() Products {
	return &productImpl{}
}
