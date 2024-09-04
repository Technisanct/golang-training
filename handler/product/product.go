package product

import (
	"github.com/gin-gonic/gin"
	"golang-training/libs/logger"
	"golang-training/logic/product"
	"net/http"
)

type handler struct {
	product product.Products
}

const logTag = "handler.product"

func (h handler) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	req := CreateProduct{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to map request body")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.product.Create(ctx, &product.CreateProductRequest{
		ProductName: req.ProductName,
		Price:       req.Price,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create product")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "product created")
}
