package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func init() {
	_ = gotenv.Load()
	viper.AutomaticEnv()

	if viper.GetString("ENV") == "" {
		panic("ENV not set")
	}

	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("URL", "localhost:3000")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_NAME", "rose_park_development")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("AUTH0_API_AUDIENCE", "")
	viper.SetDefault("AUTH0_JWK", "")

	if viper.GetString("ENV") != "production" {
		viper.AddConfigPath("./config/")
		viper.AddConfigPath("../config/")
		viper.SetConfigName(viper.GetString("ENV"))
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}

	if viper.GetString("GIN_MODE") == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
