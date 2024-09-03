package user

import (
	"context"
	"fmt"
)

func (u userImpl) Create(ctx context.Context, request *CreateUserRequest) error {
	// TODO create user
	fmt.Println("user created successfully")

	return nil
}
