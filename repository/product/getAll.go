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

func (p productImpl) GetAll(c ctx.Context) ([]model.Product, error) {
	log := logger.FromContextWithTag(c, logTag)

	newCtx, cancel := context.WithTimeout(c, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	opts := options.Find()
	_, err := p.collection.Find(newCtx, opts)
	if err != nil {
		log.Error().Err(err).
			Msg("error while retrieving products from database")

		return nil, err
	}

	return []model.Product{}, nil
	// for ptr.Next() {

	// }

	// return products, nil
}
