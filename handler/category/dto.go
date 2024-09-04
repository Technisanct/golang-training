package category

type CreateCategoryRequest struct {
	CategoryName string `json:"category_name"`
	CategoryId   int    `json:"category_id"`
}
