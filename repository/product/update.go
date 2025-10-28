package product

import (
	ctx "context"
	"errors"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (p repo) Update(c ctx.Context, productId string, updateProductRequest *model.Product) error {
	log := logger.FromContextWithTag(c, logTag)
	newCtx, cancel := ctx.WithTimeout(c, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	productID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		log.Error().Err(err).Msg("failed to update")
		return err
	}

	opts := options.Update()
	updateDoc := bson.M{"price": updateProductRequest.Price, "name": updateProductRequest.Name, "discountedPrice": updateProductRequest.DiscountedPrice, "updatedAt": time.Now()}
	update := bson.D{{Key: "$set", Value: updateDoc}}

	result, err := p.collection.UpdateOne(newCtx, bson.M{"_id": productID}, update, opts)
	if err != nil {
		log.Error().Err(err).Msg("failed to update")
		return err
	}

	if result.MatchedCount == 0 {
		log.Warn().Err(err).Msg("product not found")
		return errors.New("no record has been found")
	}

	if result.ModifiedCount == 0 {
		log.Warn().Err(err).Msg("matched product not updated")
		return errors.New("no record has been modified")
	}

	return nil
}
