package product

import "golang-training/repository/model"

// create product dto
type CreateProductResponse struct {
	Message string `json:"message"`
}
type CreateProductRequest struct {
	Name            string  `json:"name" binding:"required"`
	Price           float32 `json:"price" binding:"required,gt=0"`
	DiscountedPrice float32 `json:"discount_price" binding:"required,gt=0"`
}

// list product dto
type ListProductResponse struct {
	Message string          `json:"message"`
	Data    []model.Product `json:"data"`
}
