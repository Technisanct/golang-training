package category

import (
	"context"
	"github.com/google/uuid"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"time"
)

func (u Impl) Create(ctx context.Context, request *CreateCategoryRequest) error {
	log := logger.FromContextWithTag(ctx, logTag)

	categoryUUID := uuid.New().String()
	err := u.categoryRepo.Create(ctx, &model.Category{
		UUID:      categoryUUID,
		Name:      request.Name,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create category")
		return err
	}

	return nil
}
