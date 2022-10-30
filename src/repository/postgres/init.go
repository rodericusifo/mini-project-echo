package postgres_repository

import (
	user_postgres_repository "github.com/rodericusifo/mini-project-echo/src/repository/postgres/user"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	UserRepository user_postgres_repository.IUserRepository
}

func InitPostgresRepository(db *gorm.DB) *PostgresRepository {
	var (
		userRepository = user_postgres_repository.InitUserRepository(db)
	)

	return &PostgresRepository{
		UserRepository: userRepository,
	}
}
