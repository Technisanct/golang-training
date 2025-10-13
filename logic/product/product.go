package product

import (
	"context"
	"golang-training/config"
	"golang-training/repository/model"
	"golang-training/repository/product"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	logTag = "logic.product"
)

type productImpl struct {
	product product.Product
}
type Products interface {
	CreateProduct(c context.Context, payloadData *model.Product) error
}

// type Products interface {
// 	CreateProduct(c context.Context, payloadData *model.Product) error
// }

// type productImpl struct {
// 	product product.Product
// }

func New() *productImpl {

	var database *mongo.Database

	database = config.Get().Database.MongoDB.Client.Database(config.Get().Database.MongoDB.DBName)
	return &productImpl{
		product: product.New(database),
	}

}
