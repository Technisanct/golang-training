package product

import (
	"fmt"
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
		Name:  req.ProductName,
		Price: req.Price,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create product")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "product created")
}

func (h handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)
	fmt.Println("line 45")
	req := GetProductReq{}
	if err := c.ShouldBindUri(&req); err != nil {
		log.Error().Err(err).Msg("failed to map request body")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Println("line 52")

	response, err := h.product.Get(ctx, req.UUID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get product")
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println("line 60")

	c.JSON(http.StatusOK, mapLogicToHandler(response))
}

func (h handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	req := GetProductReq{}
	if err := c.ShouldBindUri(&req); err != nil {
		log.Error().Err(err).Msg("failed to map request body")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.product.Delete(ctx, req.UUID)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete product")
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

func mapLogicToHandler(product *product.Product) *Product {
	response := &Product{
		ID:              product.ID,
		UUID:            product.UUID,
		ProductName:     product.Name,
		Price:           product.Price,
		DiscountedPrice: product.DiscountedPrice,
		UpdatedAt:       product.UpdatedAt,
	}

	return response
}
