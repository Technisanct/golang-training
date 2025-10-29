package product

import (
	"context"
	ctx "context"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	product model.Product
)

func (p repo) Get(c context.Context, productId string) (*model.Product, error) {
	log := logger.FromContextWithTag(c, logTag)
	newCtx, cancel := ctx.WithTimeout(c, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	productID, perr := primitive.ObjectIDFromHex(productId)
	if perr != nil {
		log.Info().
			Err(perr).
			Msg("no documents found")
		return nil, perr
	}

	opts := options.FindOne()
	err := p.collection.FindOne(newCtx, bson.M{"_id": productID}, opts).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Info().
				Err(err).
				Msg("no documents found")
			return nil, err
		}
		log.Error().
			Err(err).
			Msg("failed to get product with given id")

		return nil, err
	}

	return &product, nil

}
