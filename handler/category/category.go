package category

import (
	"github.com/gin-gonic/gin"
	"golang-training/libs/logger"
	"golang-training/logic/category"
	"net/http"
)

type handler struct {
	category category.Categories
}

const logTag = "handler.category"

func (h handler) CreateCategory(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	req := CreateCategoryRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to map request body")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.category.Create(ctx, &category.CreateCategoryRequest{
		Name: req.Name,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create category")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "category created successfully")
}

func (h handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	req := GetCategoryRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		log.Error().Err(err).Msg("failed to map path body")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := h.category.Get(ctx, req.UUID)
	if err != nil {
		log.Error().Err(err).Msg("failed to get category")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, mapLogicToHandler(response))
}

func mapLogicToHandler(category *category.Category) *GetCategoryResponse {
	response := &GetCategoryResponse{
		UUID:      category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
	}

	return response
}
