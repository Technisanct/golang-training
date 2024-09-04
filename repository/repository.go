package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/repository/model"
	"golang-training/repository/user"
)

type Repository interface {
	CreateUser(ctx context.Context, doc *model.User) error
}

type repoImpl struct {
	user user.User
}

func New(db ...*DB) Repository {
	var database *mongo.Database

	if db != nil && len(db) == 1 && db[0] != nil {
		database = db[0].userMongoDB.Database(db[0].DBName)
	}

	return &repoImpl{
		user: user.New(database),
	}
}

func (r repoImpl) CreateUser(ctx context.Context, doc *model.User) error {
	return r.user.Create(ctx, doc)
}
