package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct{}

func (svc PaymentHandler) GenerateQr(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
