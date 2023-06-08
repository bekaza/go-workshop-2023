package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, username string) error
}

type UserRepo struct {
	db *gorm.DB
}

func ProvideUserRepo(conn *gorm.DB) IUserRepository {
	return &UserRepo{db: conn}
}

func (r UserRepo) Create(ctx context.Context, username string) error {
	sql, err := r.db.DB()
	if err != nil {
		return errors.New("db error")
	}
	return sql.Ping()
}
