package database

import (
	"github.com/rodericusifo/mini-project-echo/libs/config"
	"github.com/rodericusifo/mini-project-echo/libs/constant"
	postgres_repository "github.com/rodericusifo/mini-project-echo/src/repository/postgres"
)

var (
	configApps = config.ConfigApps()
)

func InitPostgres() *postgres_repository.PostgresRepository {
	db := config.ConfigureDatabase(configApps.Database.Postgres, constant.POSTGRES)
	return postgres_repository.InitPostgresRepository(db)
}
