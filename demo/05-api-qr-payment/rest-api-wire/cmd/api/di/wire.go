//go:build wireinject
// +build wireinject

package di

import (
	"example/apiwire/cmd/api/handler"
	userRepo "example/apiwire/internal/repository/user"
	"example/apiwire/internal/services/user"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeAPI(dbConn *gorm.DB) handler.Handler {
	wire.Build(DBSet, MainBindingSet, HandlerSet)
	return handler.Handler{}
}

var MainBindingSet = wire.NewSet(
	user.UserServiceSet,
)

var DBSet = wire.NewSet(
	userRepo.UserDBSet,
)

var HandlerSet = wire.NewSet(
	handler.UserHandlerSet,
	handler.HandlerSet,
)
