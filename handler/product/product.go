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
		Data:    toHandlerProductMapping(products),
	})
}

func toHandlerProductMapping(input []contract.Product) []Product {
	result := make([]Product, len(input))

	for idx, product := range input {
		result[idx] = Product{
			ID:              product.ID,
			Name:            product.Name,
			Price:           product.Price,
			DiscountedPrice: product.DiscountedPrice,
			CreatedAt:       product.CreatedAt,
			UpdatedAt:       product.UpdatedAt,
		}
	}

	return result
}
