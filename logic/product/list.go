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

	return mapRepoToLogic(products), nil
}

func mapRepoToLogic(input []*model.Product) []*contract.Product {
	results := make([]*contract.Product, 0, len(input))
	for _, product := range input {
		if product != nil {
			result := &contract.Product{
				ID:              product.ID.Hex(),
				Name:            product.Name,
				Price:           product.Price,
				DiscountedPrice: product.DiscountedPrice,
				CreatedAt:       product.CreatedAt,
				UpdatedAt:       product.UpdatedAt,
			}
			results = append(results, result)
		}
	}

	return results
}
