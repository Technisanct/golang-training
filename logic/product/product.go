package product

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/config"
	"golang-training/repository/product"
)

const (
	logTag = "logic.product"
)

type Products interface {
	Create(ctx context.Context, request *CreateProductRequest) error
	Get(ctx context.Context, uuid string) (*Product, error)
	Delete(ctx context.Context, uuid string) error
}

type productImpl struct {
	product product.Product
}

func New() Products {
	var database *mongo.Database
	database = config.Get().Database.MongoDB.Client.Database(config.Get().Database.MongoDB.DBName)
	return &productImpl{
		product: product.New(database),
	}
}
