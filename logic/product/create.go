package product

import (
	ctx "context"
	"golang-training/libs/logger"
	"golang-training/repository/model"
)

func (ps *productImpl) CreateProduct(c ctx.Context, payLoad *model.Product) error {

	log := logger.FromContextWithTag(c, logTag)

	if err := ps.product.Create(c, payLoad); err != nil {
		log.Error().Err(err).Msg("failed to create product")
		return err
	}

	return nil
}
