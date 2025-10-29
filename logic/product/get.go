package product

import (
	"context"
	"golang-training/libs/logger"
	"golang-training/logic/product/contract"
	"golang-training/repository/model"
)

func (p productImpl) Get(c context.Context, productId string) (*contract.Product, error) {
	log := logger.FromContextWithTag(c, logTag)

	product, err := p.repo.Get(c, productId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get product with given id")
		return nil, err
	}

	return mapSingleProductFromRepoToLogic(product), nil
}

func mapSingleProductFromRepoToLogic(input *model.Product) *contract.Product {
	return &contract.Product{
		ID:              input.ID.Hex(),
		UUID:            input.UUID,
		Name:            input.Name,
		Price:           input.Price,
		DiscountedPrice: input.DiscountedPrice,
		CreatedAt:       input.CreatedAt,
		UpdatedAt:       input.UpdatedAt,
	}
}
