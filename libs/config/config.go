package config

import (
	"errors"
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/rodericusifo/mini-project-echo/libs/constant"
	"github.com/rodericusifo/mini-project-echo/src/model"
)

func ConfigApps() *DefaultConfig {
	var (
		environment = flag.String("env", "", "input the environment type")
	)

	flag.Parse()

	switch constant.EnvironmentType(*environment) {
	case constant.DEV:
		viper.SetConfigFile("./environments/dev.application.yaml")
	case constant.STAG:
		viper.SetConfigFile("./environments/stag.application.yaml")
	case constant.PROD:
		viper.SetConfigFile("./environments/prod.application.yaml")
	case constant.TEST:
		viper.SetConfigFile("./environments/test.application.yaml")
	default:
		panic(errors.New("input environment type [ dev | stag | prod | test]"))
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	var conf DefaultConfig
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Panic(err)
	}

	return &conf
}

func ConfigureDatabase(ds Datasource, dialect constant.DialectDatabaseType) *gorm.DB {
	var (
		db *gorm.DB
		err error
	)
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require search_path=%s",
		ds.Url,
		ds.Username,
		ds.Password,
		ds.DatabaseName,
		ds.Port,
		ds.Schema)

	cfg := &gorm.Config{
		Logger: logger.Default,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   ds.Schema,
			SingularTable: false,
		},
	}

	if ds.DebugMode {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}

	switch dialect {
	case constant.POSTGRES:
		db, err = gorm.Open(postgres.Open(dsn), cfg)
		if err != nil {
			logrus.Panic(err)
		}
	}

	// Auto Migration Models
	db.AutoMigrate(&model.User{})

	sqlDb, err := db.DB()
	if err != nil {
		logrus.Panic(err)
	}

	sqlDb.SetConnMaxIdleTime(ds.ConnectionTimeout)
	sqlDb.SetMaxIdleConns(ds.MaxIdleConnection)
	sqlDb.SetMaxOpenConns(ds.MaxOpenConnection)

	return db
}
