package product

import (
	"context"
	"fmt"
	"golang-training/repository/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product interface {
	Create(c context.Context, doc *model.Product) error
	List(c context.Context) ([]*model.Product, error)
}

const (
	logTag         = "repository.product"
	collectionName = "products"
	KeyObjectID    = "_id"
)

type repo struct {
	collection *mongo.Collection
}

func New(db *mongo.Database) *repo {
	collection := db.Collection(collectionName)
	mod := createIndexes()

	_, err := collection.Indexes().CreateMany(context.Background(), *mod)
	if err != nil {
		panic(fmt.Sprintf("index creation failed in project repo with ERR: %v", err))
	}

	return &repo{
		collection,
	}
}

func createIndexes() *[]mongo.IndexModel {
	indexModels := &[]mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "uuid", Value: 1},
			},
			Options: options.Index(),
		},
	}
	return indexModels
}
