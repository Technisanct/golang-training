package handler

import (
	"golang-training/handler/ping"
	"golang-training/handler/product"
	"golang-training/handler/user"

	"github.com/gin-gonic/gin"
)

// InitPublicRoutes ... routes without auth
func InitPublicRoutes(router *gin.Engine) {
	ping.AddRoute(router, "/ping")
	user.AddRoute(router, "/user")
	product.AddRoute(router, "/product")
}
