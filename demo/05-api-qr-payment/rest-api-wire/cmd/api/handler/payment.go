package handler

import (
	"example/apiwire/cmd/api/handler/request"
	"example/apiwire/internal/services/payment"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var PaymentHandlerSet = wire.NewSet(ProvidePaymentHandler)

type PaymentHandler struct {
	PaymentSvc payment.PaymentService
}

func ProvidePaymentHandler(paymentService payment.PaymentService) PaymentHandler {
	return PaymentHandler{
		PaymentSvc: paymentService,
	}
}

func (svc PaymentHandler) GenerateQr(c *gin.Context) {
	var request request.GenerateQrRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
	}
	c.Status(http.StatusNoContent)
}
