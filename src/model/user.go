package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/rodericusifo/mini-project-echo/libs/types"
	"gorm.io/gorm"
)

type User struct {
	ID          string           `json:"id" gorm:"primaryKey"`
	Name        string           `json:"name"`
	Email       string           `json:"email"`
	Description types.NullString `json:"description"`
	Age         types.NullInt64  `json:"age"`
	Birthday    types.NullTime   `json:"birthday"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
	DeletedAt   gorm.DeletedAt   `json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
