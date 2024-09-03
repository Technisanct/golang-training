package user

import (
	"github.com/gin-gonic/gin"
	"golang-training/libs/logger"
	"golang-training/logic/user"
	"net/http"
)

type handler struct {
	user user.Users
}

const logTag = "handler.user"

func (h handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	req := CreateUserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error().Err(err).Msg("failed to map request body")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	err := h.user.Create(ctx, &user.CreateUserRequest{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Phone:     req.Phone,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "user created successfully")
}
