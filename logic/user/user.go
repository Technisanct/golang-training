package user

import "context"

type Users interface {
	Create(ctx context.Context, request *CreateUserRequest) error
}

type userImpl struct{}

func New() Users {
	return &userImpl{}
}
