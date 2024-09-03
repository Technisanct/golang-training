package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct{}

// Pong ... test route
func (h *handler) Pong(c *gin.Context) {
	res := &GetMessageResponse{
		Message: "pong",
	}

	c.JSON(http.StatusOK, res)
}
