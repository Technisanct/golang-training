package category

import (
	"github.com/gin-gonic/gin"
	"golang-training/libs/logger"
	"golang-training/logic/category"
	"net/http"
)

type handler struct {
	category category.Category
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
		CategoryName: req.CategoryName,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create category")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "category created successfully")
}
