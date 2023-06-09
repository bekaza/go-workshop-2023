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

func (svc UserHandler) RegisterUserHandler(c *gin.Context) {
	if err := svc.UserService.CreateUser(c.Request.Context(), "superAdmin"); err != nil {
		logrus.WithError(err).Error("cannot register user")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "internal error"})
	}
	c.JSON(http.StatusCreated, gin.H{"Msg": "hi"})
}
