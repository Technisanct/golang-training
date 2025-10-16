package product

import (
	ctx "context"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"golang-training/storage/mongodb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	products []model.Product
	product  model.Product
)

const (
	BATCH_SIZE = 100
)

func (p repo) List(c ctx.Context) ([]model.Product, error) {
	log := logger.FromContextWithTag(c, logTag)
	newCtx, cancel := ctx.WithTimeout(c, mongodb.ConnectionTimeout*time.Second)
	defer cancel()

	// start := time.Now()
	opts := options.Find().SetBatchSize(BATCH_SIZE)

	cur, err := p.collection.Find(newCtx, bson.M{}, opts)
	if err != nil {
		log.Error().Err(err).
			Msg("error while inserting data in mongodb")
		return nil, err
	}
	defer cur.Close(newCtx)

	for cur.Next(newCtx) {
		_ = cur.Decode(&product)
		products = append(products, product)
	}

	// if err = cur.All(newCtx, &products); err != nil {
	// 	log.Error().Err(err).
	// 		Msg(err.Error())
	// 	return nil, err
	// }

	// end := time.Since(start)
	// fmt.Println("total time taken to fetch products ", end)
	return products, nil
}
