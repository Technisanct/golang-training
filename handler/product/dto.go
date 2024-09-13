package product

import "time"

type CreateProduct struct {
	ProductName string  `json:"productName"`
	Price       float32 `json:"price"`
}

type GetProductReq struct {
	UUID string `uri:"uuid"`
}

type Product struct {
	ID              string    `json:"id"`
	UUID            string    `json:"uuid"`
	ProductName     string    `json:"productName"`
	Price           float32   `json:"price"`
	DiscountedPrice float32   `json:"discount"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type CreateProductResponse struct {
	Message string `json:"message"`
}
