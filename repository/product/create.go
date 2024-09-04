package product

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"
)

func (p productImpl) Create(ctx context.Context, doc *model.Product) error {
	log := logger.FromContextWithTag(ctx, logTag)
	newCtx, cancel := context.WithTimeout(ctx, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	opts := options.InsertOne()

	_, err := p.collection.InsertOne(newCtx, doc, opts)
	if err != nil {
		log.Error().Err(err).
			Interface("model", doc).
			Msg("Error inserting data in mongoDB")
		return err
	}

	return nil
}
