package config

import (
	"fmt"
	"github.com/derhabicht/rose-park-api/pkg/config"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
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

func setDefaults() {
	viper.SetDefault("AWS_ACCESS_KEY", "")
	viper.SetDefault("AWS_LOGS_REGION", "")
	viper.SetDefault("AWS_SECRET_KEY", "")
	viper.SetDefault("CLOUDWATCH_GROUP_NAME", "")
	viper.SetDefault("CLOUDWATCH_STREAM_NAME", "")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "rose_park_development")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("ENVIRONMENT", config.DEVELOPMENT)
	viper.SetDefault("MIGRATIONS_DIRECTORY", "deployments/migrations")
	viper.SetDefault("MIGRATIONS_REPO", "DerHabicht/rose-park-api")
	viper.SetDefault("URL", "localhost:3000")
}

func loadEnvironment() {
	_ = gotenv.Load()

	viper.SetEnvPrefix("RP")
	viper.AutomaticEnv()

	envStr, set := os.LookupEnv("ENVIRONMENT")
	if set {
		env, err := config.ParseEnvironment(envStr)
		if err != nil {
			panic(err)
		}
		viper.Set("ENVIRONMENT", env)
	}
}

// GetDatabaseConnectionString builds Gorm's connection string out of the configuration.
func GetDatabaseConnectionString() string {
	return fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		)
}

// GetEnvironment fetches the currently configured Environment enum value.
func GetEnvironment() config.Environment {
	env, ok := viper.Get("ENVIRONMENT").(config.Environment)
	if !ok {
		panic("api environment was not correctly set")
	}

	return env
}

// GetVersionString returns the SemVer-compliant version string of this build of the API.
func GetVersionString() string {
	return fmt.Sprintf(
		"%s+%s.%s.%s",
		BaseVersion,
		GitBranch,
		GitRevision,
		BuildTime,
	)
}

func SetUp() {
	setDefaults()
	loadEnvironment()
}