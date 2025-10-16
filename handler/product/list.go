package product

import (
	"golang-training/libs/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) ListProduct(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	products, err := h.product.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch products")
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &ListProductResponse{
		Message: "successful",
		Data:    products,
	})
}
