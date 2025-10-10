package product

import (
	ctx "context"
	"golang-training/repository/model"
)

func (p productImpl) GetOne(c ctx.Context, prodId int) (model.Product, error) {
	// log := logger.FromContextWithTag(c, logTag)

	// newCtx, cancel := context.WithTimeout(c, mongodb.ConnectionTimeout*time.Second)
	// defer cancel()

	// var product model.Product

	// opts := options.FindOne()
	// // err := p.collection.FindOne(newCtx, bson.M{ "_id": prodId }).Decode(&product)
	// // if err != nil {
	// // 	log.Error().Err(err).
	// // 		Msg("error while retrieving products from database")

	// // 	return model.Product{}, err
	// // }

	// // return nil
	return model.Product{}, nil
}
