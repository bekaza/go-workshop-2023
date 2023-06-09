package user

import (
	"context"
	"errors"
	"example/apiwire/internal/utils/timeutils"

	"github.com/google/wire"
	"gorm.io/gorm"
)

var UserDBSet = wire.NewSet(ProvideUserRepo, wire.Bind(new(IUserRepository), new(*UserRepo)))

//go:generate mockgen -source=./user.go -destination=./mock_user/mock_user_repo.go -package=mock_user_repo
type IUserRepository interface {
	Create(ctx context.Context, name string) error
}

type UserRepo struct {
	db *gorm.DB
}

func ProvideUserRepo(conn *gorm.DB) *UserRepo {
	return &UserRepo{db: conn}
}

func (r UserRepo) Create(ctx context.Context, name string) error {
	err := r.db.Debug().Create(&UserModel{
		Name:      name,
		CreatedAt: timeutils.Now(),
	}).Error
	if err != nil {
		return errors.New("internal error")
	}
	return nil
}
