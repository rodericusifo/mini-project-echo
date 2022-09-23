package repository_user_v1

import (
	"github.com/rodericusifo/mini-project-echo/libs/types"
	"github.com/rodericusifo/mini-project-echo/src/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(payload *model.User) error
	GetUser(query *types.Query, payload *model.User) (*model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
