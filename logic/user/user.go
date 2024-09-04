package user

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/config"
	"golang-training/repository/user"
)

const (
	logTag = "logic.user"
)

type Users interface {
	Create(ctx context.Context, request *CreateUserRequest) error
}

type userImpl struct {
	user user.User
}

func New() Users {
	var database *mongo.Database
	database = config.Get().Database.MongoDB.Client.Database(config.Get().Database.MongoDB.DBName)
	return &userImpl{
		user: user.New(database),
	}
}
