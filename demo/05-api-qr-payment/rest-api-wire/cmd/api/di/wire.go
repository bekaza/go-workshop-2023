//go:build wireinject
// +build wireinject

package di

import (
	"example/apiwire/cmd/api/handler"
	"example/apiwire/internal/config"
	"example/apiwire/internal/repository"
	userRepo "example/apiwire/internal/repository/user"
	"example/apiwire/internal/services/user"

	"github.com/google/wire"
)

func InitializeAPI(config config.AppConfig) (handler.Handler, func()) {
	wire.Build(DBSet, MainBindingSet, HandlerSet)
	return handler.Handler{}, nil
}

var MainBindingSet = wire.NewSet(
	user.UserServiceSet,
)

var DBSet = wire.NewSet(
	repository.PostgresDBSet,
	userRepo.UserDBSet,
)

var HandlerSet = wire.NewSet(
	handler.UserHandlerSet,
	handler.HandlerSet,
)
