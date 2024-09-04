package handler

import (
	"github.com/gin-gonic/gin"
	"golang-training/handler/category"
	"golang-training/handler/ping"
	"golang-training/handler/product"
	"golang-training/handler/user"
)

// InitPublicRoutes ... routes without auth
func InitPublicRoutes(router *gin.Engine) {
	ping.AddRoute(router, "/ping")
	user.AddRoute(router, "/user")
	category.AddRoute(router, "/category")
	product.AddRoute(router, "/product")
}
