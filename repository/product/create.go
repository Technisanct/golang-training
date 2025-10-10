package product

import (
	"context"
	ctx "context"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p productImpl) Create(c ctx.Context, payloadData *model.Product) error {

	// logs
	log := logger.FromContextWithTag(c, logTag)

	ch := make(chan error)

	newCtx, cancel := ctx.WithTimeout(c, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	go func(c context.Context, errs chan<- error) {
		opts := options.InsertOne()
		_, err := p.collection.InsertOne(newCtx, payloadData, opts)
		if err != nil {
			log.Error().Err(err).
				Interface("model", payloadData).
				Msg("error while inserting data in mongodb")
			errs <- err
		} else {
			errs <- nil
		}
	}(newCtx, ch)

	if errs := <-ch; errs != nil {
		return errs
	}

	return nil
}
