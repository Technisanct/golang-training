package product

import "time"

type CreateProductRequest struct {
	Name  string
	Price float32
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
