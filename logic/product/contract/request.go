package contract

type CreateProductRequest struct {
	Name            string
	Price           float32
	DiscountedPrice float32
}
type UpdateProductRequest struct {
	Name            string
	Price           float32
	DiscountedPrice float32
}
