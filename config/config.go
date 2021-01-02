package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	logrusCW "github.com/kdar/logrus-cloudwatchlogs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
)

var Log *logrus.Logger

func initLogging() {
	if viper.GetString("GIN_MODE") == "debug" {
		Log := logrus.StandardLogger()
		Log.SetLevel(logrus.DebugLevel)
	} else {
		sess, err := session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Config: aws.Config{Region: aws.String(viper.GetString("AWS_LOGS_REGION"))},
		})
		if err != nil {
			panic(fmt.Sprintf("Failed to establish AWS Session: %v", err))
		}

		hook, err := logrusCW.NewHook(
			viper.GetString("CLOUDWATCH_GROUP_NAME"),
			viper.GetString("CLOUDWATCH_STREAM_NAME"),
			sess,
		)
		if err != nil {
			panic(fmt.Sprintf("Failed to initialize CloudWatch logs: %v", err))
		}

		Log = logrus.New()
		Log.Hooks.Add(hook)
		Log.Out = os.Stdout
		Log.Formatter = logrusCW.NewProdFormatter()
	}
}

func init() {
	_ = gotenv.Load()
	viper.AutomaticEnv()

	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("URL", "localhost:3000")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_NAME", "rose_park_development")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("AUTH0_API_AUDIENCE", "")
	viper.SetDefault("AUTH0_JWK", "")
	viper.SetDefault("AWS_ACCESS_KEY", "")
	viper.SetDefault("AWS_SECRET_KEY", "")
	viper.SetDefault("AWS_LOGS_REGION", "")
	viper.SetDefault("CLOUDWATCH_GROUP_NAME", "")
	viper.SetDefault("CLOUDWATCH_STREAM_NAME", "")

	initLogging()
}
