package category

import (
	"context"
	"golang-training/libs/logger"
)

func (c Impl) Delete(ctx context.Context, uuid string) error {
	log := logger.FromContext(ctx)

	err := c.categoryRepo.Delete(ctx, uuid)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete category")
		return err
	}

	return nil
}
