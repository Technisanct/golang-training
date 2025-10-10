package product

import "golang-training/logic/product"

type productHandler struct {
	product product.Products
}

func NewProductHandler() *productHandler {
	return &productHandler{
		product: product.New(),
	}
}
