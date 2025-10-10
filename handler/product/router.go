package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoute(router *gin.Engine, relativePath string) {
	h := NewProductHandler()

	router.GET(relativePath, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Health": "OK",
		})
	})

	router.POST(relativePath+"/create", h.CreateProductHandler)
}
