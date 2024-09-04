package category

type CreateCategoryRequest struct {
	CategoryName string `json:"categoryName"`
	CategoryId   int    `json:"categoryId"`
}
