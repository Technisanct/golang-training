package category

import (
	"context"
)

type Category interface {
	Create(ctx context.Context, request *CreateCategoryRequest) error
}

type Impl struct{}

func New() Category {
	return &Impl{}
}
