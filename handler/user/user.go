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
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &CreateUserResponse{Message: "user created successfully"})
}

func (h handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	log := logger.FromContextWithTag(ctx, logTag)

	req := GetUserRequest{}
	if err := c.ShouldBindUri(&req); err != nil {
		log.Error().Err(err).Msg("failed to map path params")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := h.user.Get(ctx, req.UUID)
	if err != nil {
		log.Error().Err(err).Msg("failed to create user")
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, mapLogicToHandler(response))
}

func mapLogicToHandler(user *user.User) *User {
	response := &User{
		ID:        user.ID,
		UUID:      user.UUID,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	return response
}
