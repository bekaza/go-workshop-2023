package handler

import "github.com/google/wire"

var HandlerSet = wire.NewSet(NewHandler)

type Handler struct {
	UserHandler    UserHandler
	PaymentHandler PaymentHandler
}

func NewHandler(userHandler UserHandler, paymentHandler PaymentHandler) Handler {
	return Handler{
		UserHandler:    userHandler,
		PaymentHandler: paymentHandler,
	}
}
