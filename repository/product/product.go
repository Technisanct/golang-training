package product

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-training/repository/model"
)

type Product interface {
	Create(ctx context.Context, doc *model.Product) error
	Find(ctx context.Context, uuid string) (*model.Product, error)
	DeleteOne(ctx context.Context, uuid string) error
}

const (
	logTag = "repository.product"

	collectionName = "product"
	KeyObjectID    = "_id"
)

type productImpl struct {
	collection *mongo.Collection
}

func New(database *mongo.Database) Product {
	collection := database.Collection(collectionName)

	mod := createIndexes()

	_, err := collection.Indexes().CreateMany(context.Background(), *mod)
	if err != nil {
		panic(fmt.Sprintf("index creation failed in the project with ERP: %v", err))
	}

	return &productImpl{
		collection: collection,
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
