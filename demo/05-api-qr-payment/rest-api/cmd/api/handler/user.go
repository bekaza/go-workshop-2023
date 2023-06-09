package handler

import (
	"example/restapi/internal/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserService user.UserService
}

func (svc UserHandler) RegisterUserHandler(c *gin.Context) {
	if err := svc.UserService.CreateUser(c.Request.Context(), "superAdmin"); err != nil {
		logrus.WithError(err).Error("cannot register user")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "internal error"})
	}
	c.JSON(http.StatusCreated, gin.H{"Msg": "hi"})
}
