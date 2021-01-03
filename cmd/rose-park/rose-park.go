package main

import (
	"fmt"
	"github.com/derhabicht/rose-park-api/internal/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	instance "github.com/derhabicht/rose-park-api/internal/config"
	"github.com/derhabicht/rose-park-api/pkg/config"
	"github.com/derhabicht/rose-park-api/pkg/data"
)

// BaseVersion is the SemVer-formatted base version for this build of the API.
// GetVersionString will concat this with BuildTime, GitRevision, and GitBranch to return the full version string.
const BaseVersion = "0.1.0-develop"

// BuildTime is a timestamp of the build.
// GetVersionString will concat this with BaseVersion, GitRevision, and GitBranch to return the full version string.
// This variable is set at compile-time.
var BuildTime string

// GitRevision is the current Git commit ID.
// If the tree is dirty at compile-time, an "x-" is prepended to the hash.
// GetVersionString will concat this with BaseVersion, BuildTime, and GitBranch to return the full version string.
// This variable is set at compile-time.
var GitRevision string

// GitBranch is the name of the active Git branch at compile-time.
// GetVersionString will concat this with BaseVersion, BuildTime, and GitRevision to return the full version string.
// This variable is set at compile-time.
var GitBranch string

func getVersionString() string {
	return fmt.Sprintf(
		"%s+%s.%s.%s",
		BaseVersion,
		GitBranch,
		GitRevision,
		BuildTime,
	)
}

func migrateDatabase(logger *logrus.Entry, db data.IConnection) {
	var envBranch string
	switch instance.GetEnvironment() {
	case config.PRODUCTION:
		envBranch = "master"
	case config.STAGE:
		envBranch = "develop"
	default:
		logger.Info("Current environment is neither production nor stage. Skipping migrations.")
		return
	}

	err := db.Migrate(
		viper.GetString("MIGRATIONS_REPO"),
		viper.GetString("MIGRATIONS_DIRECTORY"),
		envBranch,
	)

	if err != nil {
		if err.Error() == "no change" {
			logger.Info("Database already at latest version.")
		} else {
			logger.Errorf("Failed to migrate database: %v", err)
		}
	}
}

func initConfig() *logrus.Logger {
	instance.SetUp()
	logger, err := instance.SetupLogging()
	if err != nil {
		panic(err)
	}

	return logger
}

func initDatabase(logger *logrus.Logger) data.IConnection {
	l := logger.WithFields(logrus.Fields{
		"database_name": viper.GetString("DB_NAME"),
	})

	l.Info("Attempting to connect to database")
	db, err := data.NewConnection(instance.GetDatabaseConnectionString())
	if err != nil {
		logger.WithFields(logrus.Fields{
			"database_name": viper.GetString("DB_NAME"),
		}).Panic("Failed to connect to database")
	}

	migrateDatabase(l, db)

	return db
}

func main() {
	// Initialize server configuration
	logger := initConfig()
	logger.WithFields(logrus.Fields{
		"version": getVersionString(),
	}).Info("Server configuration loaded")

	// Connect to database
	db := initDatabase(logger)
	logger.Info("Database initialized")

	// Configure middleware

	// Configure router
	r := router.SetUpRouter(db, logger, getVersionString())

	// Start server
	_ = r.Run(viper.GetString("url"))
}