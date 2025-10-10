package user

import (
	"context"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u userImpl) Create(ctx context.Context, doc *model.User) error {
	log := logger.FromContextWithTag(ctx, logTag)
	newCtx, cancel := context.WithTimeout(ctx, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	opts := options.InsertOne()

	_, err := u.collection.InsertOne(newCtx, doc, opts)
	if err != nil {
		log.Error().Err(err).
			Interface("model", doc).
			Msg("error while inserting data in mongodb")
		return err
	}

	return nil
}
