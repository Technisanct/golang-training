package product

type CreateProduct struct {
	ProductName string  `json:"productName"`
	ProductID   string  `json:"productID"`
	Price       float32 `json:"price"`
}
