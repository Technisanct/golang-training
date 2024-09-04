package user

import (
	"context"
	"golang-training/repository"
)

const (
	logTag = "logic.user"
)

type Users interface {
	Create(ctx context.Context, request *CreateUserRequest) error
}

type userImpl struct {
	repo repository.Repository
}

func New() Users {
	return &userImpl{
		repo: repository.New(),
	}
}
