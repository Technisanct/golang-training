package product

type (
	CreateProductResponse struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	CreateProductRequest struct {
		Name            string  `json:"name" binding:"required"`
		Price           float32 `json:"price" binding:"required,gt=0"`
		DiscountedPrice float32 `json:"discount_price" binding:"required,gt=0"`
	}
)
