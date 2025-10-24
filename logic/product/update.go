package product

import (
	"context"
	"golang-training/libs/logger"
	"golang-training/logic/product/contract"
	"golang-training/repository/model"
)

func (ps *productImpl) Update(c context.Context, productId string, updateProductRequest *contract.UpdateProductRequest) error {
	log := logger.FromContextWithTag(c, logTag)

	err := ps.repo.Update(c, productId, &model.Product{
		Name:            updateProductRequest.Name,
		Price:           updateProductRequest.Price,
		DiscountedPrice: updateProductRequest.DiscountedPrice,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to update product")
		return err
	}

	return nil
}
