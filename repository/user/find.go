package user

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"
)

func (u userImpl) Find(ctx context.Context, uuid string) (*model.User, error) {
	log := logger.FromContextWithTag(ctx, logTag)
	log.Info().Msg("exc find all in repo")

	newCtx, cancel := context.WithTimeout(ctx, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	filters := bson.D{{"uuid", uuid}}

	var result model.User
	err := u.collection.FindOne(newCtx, filters).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Info().Msg("no document found")
			return nil, nil
		}
		msg := "Failed to find data from db"
		log.Error().Err(err).Msg(msg)
		return nil, err
	}

	return &result, nil
}
