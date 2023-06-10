package repository

import (
	"example/apiwire/internal/config"
	"fmt"

	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresDBSet = wire.NewSet(NewPostgresConnection)

func NewPostgresConnection(cfg config.AppConfig) (*gorm.DB, func()) {
	config := cfg.Database
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.Username, config.Password, config.DatabaseName, config.Port)
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: DSN,
	}), &gorm.Config{})
	if err != nil {
		panic("can not connect to database")
	}

	cleanup := func() {
		sqlDB, err := conn.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
	return conn, cleanup
}
