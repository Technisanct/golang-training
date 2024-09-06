package category

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-training/repository/model"
)

type Category interface {
	Create(ctx context.Context, doc *model.Category) error
	Find(ctx context.Context, uuid string) (*model.Category, error)
	Delete(ctx context.Context, uuid string) error
}

const (
	logTag = "repository.category"

	collectionNAme = "Categories"
	keyObjectID    = "_id"
)

type Impl struct {
	collection *mongo.Collection
}

func New(database *mongo.Database) Category {
	collection := database.Collection(collectionNAme)

	mod := createIndexes()

	_, err := collection.Indexes().CreateMany(context.Background(), *mod)
	if err != nil {
		panic(fmt.Sprintf("index creation failed in project repo with ERR: %v", err))
	}

	return &Impl{
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
