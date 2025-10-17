package product

import (
	"time"
)

type CreateProductResponse struct {
	Message string `json:"message"`
}
type CreateProductRequest struct {
	Name            string  `json:"name" binding:"required"`
	Price           float32 `json:"price" binding:"required,gt=0"`
	DiscountedPrice float32 `json:"discount_price" binding:"required,gt=0"`
}

type Product struct {
	ID              string
	UUID            string
	Name            string
	Price           float32
	DiscountedPrice float32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
type ListProductResponse struct {
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}
