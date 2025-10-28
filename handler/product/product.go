package product

import (
	"errors"
	"golang-training/libs/logger"
	"golang-training/logic/product"
	"golang-training/logic/product/contract"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
		Data:    mapLogicToHandler(products),
	})
}

func mapLogicToHandler(input []*contract.Product) []Product {
	results := make([]Product, 0, len(input))

	for _, product := range input {

		if product != nil {
			result := Product{
				ID:              product.ID,
				Name:            product.Name,
				Price:           product.Price,
				DiscountedPrice: product.DiscountedPrice,
				CreatedAt:       product.CreatedAt,
				UpdatedAt:       product.UpdatedAt,
			}
			results = append(results, result)
		}
	}

	return results
}

func (h *handler) GetProduct(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)
	productId := c.Param("id")

	product, err := h.product.Get(ctx, productId)
	if err != nil {
		log.Error().Err(err).Msg("failed to get product")

		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &GetProductResponse{
		Message: "successful",
		Data:    mapSingleProductFromLogicToHandler(product),
	})
}

func mapSingleProductFromLogicToHandler(input *contract.Product) Product {
	return Product{
		ID:              input.ID,
		UUID:            input.UUID,
		Name:            input.Name,
		Price:           input.Price,
		DiscountedPrice: input.DiscountedPrice,
		CreatedAt:       input.CreatedAt,
		UpdatedAt:       input.UpdatedAt,
	}
}
