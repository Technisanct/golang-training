package category

import (
	"context"
	"fmt"
)

func (u Impl) Create(ctx context.Context, request *CreateCategoryRequest) error {

	fmt.Println("category created")
	return nil
}
