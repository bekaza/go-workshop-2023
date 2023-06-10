package router

import (
	"example/apiwire/cmd/api/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GenerateRouter(h handler.Handler) *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	router.GET("/user", h.UserHandler.RegisterUserHandler)
	router.POST("/qr", h.PaymentHandler.GenerateQr)

	return router
}
