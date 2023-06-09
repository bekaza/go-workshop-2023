package handler

import "github.com/google/wire"

var HandlerSet = wire.NewSet(NewHandler)

type Handler struct {
	UserHandler UserHandler
}

func NewHandler(userHandler UserHandler) Handler {
	return Handler{
		UserHandler: userHandler,
	}
}
