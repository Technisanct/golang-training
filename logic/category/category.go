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

type Categories interface {
	Create(ctx context.Context, request *CreateCategoryRequest) error
	Get(ctx context.Context, uuid string) (*Category, error)
	Delete(ctx context.Context, uuid string) error
}

type Impl struct {
	categoryRepo category.Category
}

func New() Categories {
	var database *mongo.Database
	database = config.Get().Database.MongoDB.Client.Database(config.Get().Database.MongoDB.DBName)
	return &Impl{
		categoryRepo: category.New(database),
	}
}
