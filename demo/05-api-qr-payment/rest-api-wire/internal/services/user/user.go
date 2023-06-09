package user

import (
	"context"
	"example/restapi/internal/repository/user"
)

type UserService interface {
	CreateUser(ctx context.Context, username string) error
}

type userServiceImpl struct {
	UserRepo user.IUserRepository
}

func ProvideUserService(userRepo user.IUserRepository) UserService {
	return &userServiceImpl{
		UserRepo: userRepo,
	}
}

func (s userServiceImpl) CreateUser(ctx context.Context, username string) error {
	return s.UserRepo.Create(ctx, username)
}
