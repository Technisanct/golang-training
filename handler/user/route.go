package user

import (
	"github.com/gin-gonic/gin"
	"golang-training/logic/user"
)

func AddRoute(router *gin.Engine, relativePath string) {
	h := handler{
		user: user.New(),
	}

	// /user
	router.POST(relativePath, h.CreateUser)
	// /user/:uuid
	router.GET(relativePath+"/:uuid", h.Get)
}
