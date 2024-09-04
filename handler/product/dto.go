package product

type CreateProduct struct {
	ProductName string  `json:"productName"`
	Price       float32 `json:"price"`
}
