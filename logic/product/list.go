package product

import (
	"context"
	"golang-training/libs/logger"
	"golang-training/repository/model"
)

func (ps *productImpl) List(c context.Context) ([]model.Product, error) {
	log := logger.FromContextWithTag(c, logTag)

	products, err := ps.repo.List(c)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch list of products")
		return nil, err
	}

	return products, nil
}
