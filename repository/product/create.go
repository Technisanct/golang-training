package product

import (
	ctx "context"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p productImpl) Create(c ctx.Context, payloadData *model.Product) error {

	log := logger.FromContextWithTag(c, logTag)
	newCtx, cancel := ctx.WithTimeout(c, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	opts := options.InsertOne()

	_, err := p.collection.InsertOne(newCtx, payloadData, opts)
	if err != nil {
		log.Error().Err(err).
			Interface("model", payloadData).
			Msg("error while inserting data in mongodb")
		return err
	}

	return nil
}
