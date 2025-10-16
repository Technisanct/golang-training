package product

import (
	"golang-training/libs/logger"
	"golang-training/logic/product"
	"golang-training/logic/product/contract"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	product product.Products
}

const logTag = "handler.product"

// create handler
func (h handler) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	req := &CreateProductRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to map request body")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.product.Create(ctx, &contract.CreateProductRequest{
		Name:            req.Name,
		Price:           req.Price,
		DiscountedPrice: req.DiscountedPrice,
	})
	if err != nil {
		log.Error().Err(err).Msg("error while creating product")
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, &CreateProductResponse{
		Message: "product successfully created",
	})
}

// list handler
func (h *handler) ListProduct(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	products, err := h.product.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to fetch products")
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &ListProductResponse{
		Message: "successful",
		Data:    products,
	})
}
