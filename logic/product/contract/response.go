package contract

import "time"

type Product struct {
	ID              string
	UUID            string
	Name            string
	Price           float32
	DiscountedPrice float32
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
