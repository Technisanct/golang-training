package category

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang-training/libs/logger"
	"golang-training/storage/mongodb"
	"time"
)

func (c Impl) Delete(ctx context.Context, uuid string) error {
	log := logger.FromContextWithTag(ctx, logTag)

	newCtx, cancel := context.WithTimeout(ctx, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	filter := bson.D{{"uuid", uuid}}

	_, err := c.collection.DeleteOne(newCtx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Info().Msg("No document found to delete")
			return err
		}
		msg := "Failed to delete document"
		log.Error().Err(err).Msg(msg)
		return err
	}

	return nil
}
