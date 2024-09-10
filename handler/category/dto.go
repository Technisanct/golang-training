package category

import "time"

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type GetCategoryRequest struct {
	UUID string `uri:"uuid"`
}

type GetCategoryResponse struct {
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type GetDeleteCategoryRequest struct {
	UUID string `uri:"uuid"`
}
