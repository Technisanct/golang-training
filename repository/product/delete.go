package product

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/libs/logger"
	"golang-training/storage/mongodb"
	"time"
)

func (p productImpl) DeleteOne(ctx context.Context, uuid string) error {
	log := logger.FromContextWithTag(ctx, logTag)
	log.Info().Msg("exc find all in repo")

	newCtx, cancel := context.WithTimeout(ctx, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	filters := bson.D{{"uuid", uuid}}

	result, err := p.collection.DeleteOne(newCtx, filters)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Info().Msg("No Documents Found")
			return err
		}
		msg := "Failed to fetch data"
		log.Error().Err(err).Msg(msg)
		return nil
	}
	if result.DeletedCount != 0 {
		log.Info().Msg("Document deleted successfully")
		return nil
	}

	return nil
}
