package product

import (
	"context"
	"github.com/google/uuid"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"time"
)

func (p productImpl) Create(ctx context.Context, request *CreateProductRequest) error {
	log := logger.FromContextWithTag(ctx, logTag)

	productUUID := uuid.New().String()
	price := request.Price

	var discountedPrice float32
	discountedPrice = price - (price * 0.05)

	err := p.product.Create(ctx, &model.Product{
		UUID:            productUUID,
		Name:            request.ProductName,
		Price:           price,
		DiscountedPrice: discountedPrice,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create product")
		return err
	}

	return nil
}
