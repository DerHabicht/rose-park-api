package database

import (
	"fmt"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	_ "github.com/derhabicht/rose-park/config"
)

var DB *gorm.DB = nil

func init() {
	var err error = nil

	connectionStr := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
	)

	// TODO: Figure out how to add defer DB.Close()
	DB, err = gorm.Open("postgres", connectionStr)
	if err != nil {
		panic(fmt.Sprintf("Unable to establish database connection: %v", err))
	}

	if viper.GetString("GIN_MODE") == "debug" {
			DB.SetLogger(&GormLogger{})
			DB.LogMode(true)
	}
}

type GormLogger struct{}

func (g *GormLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		logrus.WithFields(logrus.Fields{
			"module": "gorm",
			"type": "sql",
		}).Debug(v[3])
	}
	if v[0] == "log" {
		logrus.WithFields(logrus.Fields{
			"module": "gorm",
			"type": "log",
		}).Debug(v[2])
	}
}