package category

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"
)

func (u Impl) Create(ctx context.Context, doc *model.Category) error {
	log := logger.FromContextWithTag(ctx, logTag)
	newCtx, cancel := context.WithTimeout(ctx, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	opts := options.InsertOne()

	_, err := u.collection.InsertOne(newCtx, doc, opts)
	if err != nil {
		log.Error().Err(err).Interface("model", doc).Msg("error while inserting data in mongodb")
		return err
	}

	return nil
}
