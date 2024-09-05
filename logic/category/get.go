package category

import (
	"context"

	"golang-training/libs/logger"
	"golang-training/repository/model"
)

func (i Impl) Get(ctx context.Context, uuid string) (*Category, error) {
	log := logger.FromContextWithTag(ctx, logTag)

	category, err := i.categoryRepo.Find(ctx, uuid)
	if err != nil {
		log.Error().Err(err).Msg("failed to find category")
		return nil, err
	}

	return mapRepoToLogic(category), nil
}

func mapRepoToLogic(category *model.Category) *Category {
	response := &Category{
		ID:        category.ID.Hex(),
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}
	return response
}
