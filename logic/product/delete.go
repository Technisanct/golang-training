package product

import (
	"context"
	"golang-training/libs/logger"
)

func (p productImpl) Delete(ctx context.Context, uuid string) error {
	log := logger.FromContextWithTag(ctx, logTag)

	err := p.product.DeleteOne(ctx, uuid)
	if err != nil {
		log.Error().Err(err).Msg("delete failed")
		return err
	}

	return nil
}
