package category

import "time"

type CreateCategoryRequest struct {
	CategoryName string
	CategoryId   int
}

type Category struct {
	ID           string
	CategoryName string
	CreatedAt    time.Time
}
