package handler

import (
	"github.com/gin-gonic/gin"
	"golang-training/handler/ping"
)

// InitPublicRoutes ... routes without auth
func InitPublicRoutes(router *gin.Engine) {
	ping.AddRoute(router, "/ping")
}
