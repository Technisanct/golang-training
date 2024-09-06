package category

import (
	"github.com/gin-gonic/gin"
	"golang-training/logic/category"
)

func AddRoute(router *gin.Engine, relativePath string) {
	h := handler{
		category: category.New(),
	}

	router.POST(relativePath, h.CreateCategory)
	router.GET(relativePath+"/:uuid", h.Get)
	router.DELETE(relativePath+"/:uuid", h.Delete)

}
