package user

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-training/repository/model"
)

type User interface {
	Create(ctx context.Context, doc *model.User) error
	Find(ctx context.Context, uuid string) (*model.User, error)
}

const (
	logTag = "repository.user"

	collectionName = "users"
	KeyObjectID    = "_id"
)

type userImpl struct {
	collection *mongo.Collection
}

func New(database *mongo.Database) User {
	collection := database.Collection(collectionName)

	mod := createIndexes()

	_, err := collection.Indexes().CreateMany(context.Background(), *mod)
	if err != nil {
		panic(fmt.Sprintf("index creation failed in project repo with ERR: %v", err))
	}

	return &userImpl{
		collection: collection,
	}
}

// createIndexes... define indexes for the collection
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
