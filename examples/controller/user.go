package controller

import (
	"net/http"
	"webservice-library/examples/models"

	"github.com/gin-gonic/gin"
)

// UserController is the basic controller for /Users API
type UserController struct {
	Controller
}

func (ctrl *UserController) GetAll(ctx *gin.Context) {
	user := models.User{
		ID:   1,
		Name: "test user",
		Role: "admin",
	}

	ctx.JSON(http.StatusOK, user)
	ctx.Abort()
	return
}
