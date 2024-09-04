package product

type CreateProduct struct {
	ProductName string  `json:"product_name"`
	ProductID   string  `json:"product_id"`
	Price       float32 `json:"price"`
}
