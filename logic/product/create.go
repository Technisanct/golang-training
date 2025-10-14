package product

import (
	ctx "context"
	"golang-training/libs/logger"
	"golang-training/logic/product/contract"
	"golang-training/repository/model"
	"time"
)

func (ps *productImpl) Create(c ctx.Context, request *contract.CreateProductRequest) error {
	log := logger.FromContextWithTag(c, logTag)

	err := ps.product.Create(c, &model.Product{
		Name:            request.Name,
		Price:           request.Price,
		DiscountedPrice: request.DiscountedPrice,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create product")
		return err
	}

	return nil
}
