package product

import (
	"github.com/gin-gonic/gin"
	"golang-training/logic/product"
)

func AddRoute(router *gin.Engine, relativePath string) {
	h := handler{
		product: product.New(),
	}

	// /product
	router.POST(relativePath, h.CreateProduct)

	// /product/:UUID
	router.GET(relativePath+"/:uuid", h.Get)
}
