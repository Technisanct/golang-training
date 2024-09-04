package category

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/config"
	"golang-training/repository/category"
)

const (
	logTag = "logic.category"
)

type Category interface {
	Create(ctx context.Context, request *CreateCategoryRequest) error
}

type Impl struct {
	category category.Category
}

func New() Category {
	var database *mongo.Database
	database = config.Get().Database.MongoDB.Client.Database(config.Get().Database.MongoDB.DBName)
	return &Impl{
		category: category.New(database),
	}
}
