package user

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"column:id;primary_key"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
