package user

import (
	"context"
	"golang-training/libs/logger"
	"golang-training/repository/model"
)

func (u userImpl) Get(ctx context.Context, uuid string) (*User, error) {
	log := logger.FromContextWithTag(ctx, logTag)

	user, err := u.user.Find(ctx, uuid)
	if err != nil {
		log.Error().Err(err).Msg("failed to find user from db")
		return nil, err
	}

	return mapRepoToLogic(user), nil

}

func mapRepoToLogic(user *model.User) *User {
	response := &User{
		ID:        user.ID.Hex(),
		UUID:      user.UUID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return response
}
