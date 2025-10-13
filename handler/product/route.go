package product

import (
	"golang-training/logic/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoute(router *gin.Engine, relativePath string) {
	h := handler{
		product: product.New(),
	}

	router.GET(relativePath, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Health": "OK",
		})
	})

	router.POST(relativePath+"/create", h.CreateProductHandler)
}
