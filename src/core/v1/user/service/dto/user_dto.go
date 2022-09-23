package dto_service_user_v1

import (
	"time"
)

type UserDTO struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Description *string    `json:"description"`
	Age         *int64     `json:"age"`
	Birthday    *time.Time `json:"birthday"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-"`
}
