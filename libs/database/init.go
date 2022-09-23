package database

import (
	"github.com/rodericusifo/mini-project-echo/libs/config"

	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	configApps := config.ConfigApps()
	db := config.ConfigureDatabase(configApps.Database)
	return db
}
