package resource_user_v1

import (
	"github.com/rodericusifo/mini-project-echo/libs/database"
	"github.com/rodericusifo/mini-project-echo/libs/types"
	"github.com/rodericusifo/mini-project-echo/src/model"
	postgres_repository "github.com/rodericusifo/mini-project-echo/src/repository/postgres"
)

type IUserResource interface {
	CreateUser(payload *model.User) error
	GetUser(query *types.Query, payload *model.User) (*model.User, error)
}

type UserResource struct {
	Postgres *postgres_repository.PostgresRepository
}

func InitUserResource() IUserResource {
	var (
		postgresDatabase = database.InitPostgresDatabase()
	)

	return &UserResource{
		Postgres: postgresDatabase,
	}
}
