package product

import "time"

type CreateProductResponse struct {
	Message string `json:"message"`
}

type CreateProductRequest struct {
	Name            string  `json:"name" binding:"required"`
	Price           float32 `json:"price" binding:"required,gt=0"`
	DiscountedPrice float32 `json:"discount_price" binding:"required,gt=0"`
}

type Product struct {
	ID              string    `json:"_id,omitempty"`
	UUID            string    `json:"uuid"`
	Name            string    `json:"name"`
	Price           float32   `json:"price"`
	DiscountedPrice float32   `json:"discountedPrice"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
