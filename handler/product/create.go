package product

import (
	"fmt"
	"golang-training/libs/logger"
	"golang-training/repository/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

const logTag = "handler.user"

func (ph productHandler) CreateProductHandler(c *gin.Context) {

	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	fmt.Println("executing")

	var productBody CreateProductRequest

	if err := c.ShouldBindJSON(&productBody); err != nil {
		log.Error().Err(err).Msg("failed to map request body")
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	ph.product.CreateProduct(ctx, &model.Product{
		Name:            productBody.Name,
		Price:           productBody.Price,
		DiscountedPrice: productBody.DiscountedPrice,
	})

	c.JSON(http.StatusCreated, &CreateProductResponse{
		Status:  "success",
		Message: "product successfully created",
	})
}
