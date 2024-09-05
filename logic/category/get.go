package category

import (
	"context"

	"golang-training/libs/logger"
	"golang-training/repository/model"
)

func (u Impl) Get(ctx context.Context, uuid string) (*Category, error) {
	log := logger.FromContextWithTag(ctx, logTag)

	category, err := u.categoryRepo.Find(ctx, uuid)
	if err != nil {
		log.Error().Err(err).Msg("failed to find category")
		return nil, err
	}

	return mapRepoToLogic(category), nil
}

func mapRepoToLogic(category *model.Category) *Category {
	response := &Category{
		ID:           category.ID.Hex(),
		CategoryName: category.CategoryName,
		CreatedAt:    category.CreatedAt,
	}
	return response
}
