package product

import (
	"context"
	"golang-training/libs/logger"
	"golang-training/repository/model"
)

func (p productImpl) Get(ctx context.Context, uuid string) (*Product, error) {
	log := logger.FromContextWithTag(ctx, logTag)

	product, err := p.product.Find(ctx, uuid)
	if err != nil {
		log.Error().Err(err).Msg("get product failed")
		return nil, err
	}

	return mapRepoToLogic(product), nil
}

func mapRepoToLogic(product *model.Product) *Product {
	response := &Product{
		ID:              product.ID.Hex(),
		UUID:            product.UUID,
		Name:            product.Name,
		Price:           product.Price,
		DiscountedPrice: product.DiscountedPrice,
		UpdatedAt:       product.UpdatedAt,
	}

	return response
}
