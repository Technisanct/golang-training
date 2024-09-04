package user

import (
	"context"
	"github.com/google/uuid"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"time"
)

func (u userImpl) Create(ctx context.Context, request *CreateUserRequest) error {
	log := logger.FromContextWithTag(ctx, logTag)

	userUUID := uuid.New().String()
	err := u.repo.CreateUser(ctx, &model.User{
		UUID:      userUUID,
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Email:     request.Email,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		return err
	}

	return nil
}
