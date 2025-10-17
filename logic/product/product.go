package product

import (
	"context"
	"golang-training/config"
	"golang-training/repository/product"

	"golang-training/logic/product/contract"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	logTag = "logic.product"
)

type Products interface {
	Create(c context.Context, request *contract.CreateProductRequest) error
	List(c context.Context) ([]contract.Product, error)
}
type productImpl struct {
	repo product.Product
}

func New() Products {
	var database *mongo.Database
	database = config.Get().Database.MongoDB.Client.Database(config.Get().Database.MongoDB.DBName)
	return &productImpl{
		repo: product.New(database),
	}
}
