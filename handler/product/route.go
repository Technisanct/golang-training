package product

import (
	"golang-training/logic/product"

	"github.com/gin-gonic/gin"
)

func AddRoute(router *gin.Engine, relativePath string) {
	h := handler{
		product: product.New(),
	}

	// /product
	router.GET(relativePath, h.ListProduct)
	// /product/create
	router.POST(relativePath+"/create", h.CreateProduct)
}
