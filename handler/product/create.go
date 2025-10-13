package product

import (
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const logTag = "handler.user"

func (h handler) CreateProductHandler(c *gin.Context) {

	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	req := &CreateProductRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to map request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	err := h.product.CreateProduct(ctx, &model.Product{
		Name:            req.Name,
		Price:           req.Price,
		DiscountedPrice: req.DiscountedPrice,
	})
	if err != nil {
		log.Error().Err(err).Msg("error while creating product")
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &CreateProductResponse{
		Status:  "success",
		Message: "product successfully created",
	})
}
