package ping

import (
	"github.com/gin-gonic/gin"
)

// AddRoute ... add routes under a handler
func AddRoute(router *gin.Engine, relativePath string) {
	h := &handler{}

	// /ping
	router.GET(relativePath, h.Pong)
}
