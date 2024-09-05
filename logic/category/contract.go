package category

import "time"

type CreateCategoryRequest struct {
	Name string
	Id   int
}

type Category struct {
	ID        string
	Name      string
	CreatedAt time.Time
}
