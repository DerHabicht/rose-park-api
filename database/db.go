package database

import (
	"fmt"
	"github.com/derhabicht/rose-park/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	_ "github.com/derhabicht/rose-park/config"
)

var DB *gorm.DB = nil

func migrateDatabase() error {
	var envBranch string
	switch viper.GetString("ENV") {
	case "production":
		envBranch = "master"
	case "stage":
		envBranch = "develop"
	default:
		config.Log.WithFields(logrus.Fields{
			"environment": viper.GetString("ENV"),
		}).Info("Environment is not set to stage or production, skipping migrations.")
		return nil
	}

	migrationSource := fmt.Sprintf(
		"github://%s:%s@DerHabicht/rose-park-api/database/migrations#%s",
		viper.GetString("GITHUB_USER"),
		viper.GetString("GITHUB_ACCESS_TOKEN"),
		envBranch,
		)
	drv, err := postgres.WithInstance(DB.DB(), &postgres.Config{})
	if err != nil {
		return errors.WithMessage(err, "Failed to create migration driver")
	}
	m, err := migrate.NewWithDatabaseInstance(migrationSource, "postgres", drv)
	if err != nil {
		return errors.WithMessage(err, "Failed to create a new migration")
	}

	config.Log.WithFields(logrus.Fields{
		"db_name": viper.GetString("DB_NAME"),
	}).Info("Migrating database.")
	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			config.Log.WithFields(logrus.Fields{
				"db_name": viper.GetString("DB_NAME"),
			}).Info("Database already at latest version.")
			return nil
		}
		return errors.WithMessage(err, "Migration failed")
	}
	config.Log.WithFields(logrus.Fields{
		"db_name": viper.GetString("DB_NAME"),
	}).Info("Database migration successful.")

	return nil
}

func init() {
	var err error = nil

	connectionStr := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
	)

	config.Log.WithFields(logrus.Fields{
		"db_name": viper.GetString("DB_NAME"),
	}).Info("Connecting to database.")
	// TODO: Figure out how to add defer DB.Close()
	DB, err = gorm.Open("postgres", connectionStr)
	if err != nil {
		panic(fmt.Sprintf("Unable to establish database connection: %v", err))
	}
	config.Log.WithFields(logrus.Fields{
		"db_name": viper.GetString("DB_NAME"),
	}).Info("Database connection established.")

	err = migrateDatabase()
	if err != nil {
		config.Log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to migrate database.")
	}
}
