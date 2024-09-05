package product

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"
)

func (p productImpl) Find(ctx context.Context, uuid string) (*model.Product, error) {
	log := logger.FromContextWithTag(ctx, logTag)
	log.Info().Msg("exc find all in repo")

	newCtx, cancel := context.WithTimeout(ctx, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	filters := bson.D{{"uuid", uuid}}

	var result model.Product
	err := p.collection.FindOne(newCtx, filters).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Info().Msg("No Documents Found")
			return nil, nil
		}
		msg := "Failed to fetch data"
		log.Error().Err(err).Msg(msg)
		return nil, err
	}

	return &result, nil
}
