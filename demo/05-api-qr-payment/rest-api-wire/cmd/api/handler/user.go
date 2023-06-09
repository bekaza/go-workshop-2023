package handler

import (
	"example/apiwire/internal/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var UserHandlerSet = wire.NewSet(ProvideUserHandler)

type UserHandler struct {
	UserService user.UserService
}

func ProvideUserHandler(userService user.UserService) UserHandler {
	return UserHandler{UserService: userService}
}

type UserRequest struct {
	Name string `json:"name"`
}

func (svc UserHandler) RegisterUserHandler(c *gin.Context) {
	var req UserRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "bad request"})
		return
	}
	if err := svc.UserService.CreateUser(c.Request.Context(), req.Name); err != nil {
		logrus.WithError(err).Error("cannot register user")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "internal error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Msg": "hi"})
}
