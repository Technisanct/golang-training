package product

import (
	"context"
	"golang-training/libs/logger"
	"golang-training/logic/product/contract"
	"golang-training/repository/model"
)

func (ps *productImpl) List(c context.Context) ([]*contract.Product, error) {
	log := logger.FromContextWithTag(c, logTag)

	products, err := ps.repo.List(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch list of products")
		return nil, err
	}

	return toLogicProductMapping(products), nil
}

func toLogicProductMapping(input []*model.Product) []*contract.Product {
	result := make([]*contract.Product, len(input))
	for idx, product := range input {
		result[idx] = &contract.Product{
			ID:              product.ID.Hex(),
			Name:            product.Name,
			Price:           product.Price,
			DiscountedPrice: product.DiscountedPrice,
			CreatedAt:       product.CreatedAt,
			UpdatedAt:       product.UpdatedAt,
		}
	}
	return result
}
