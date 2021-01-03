package config

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	logrusCloudWatch "github.com/kdar/logrus-cloudwatchlogs"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/derhabicht/rose-park-api/pkg/config"
)

func setupCloudWatchLogger() (*logrus.Logger, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{Region: aws.String(viper.GetString("AWS_LOGS_REGION"))},
	})
	if err != nil {
		return nil, errors.Errorf("failed to establish an AWS session: %v", err)
	}

	hook, err := logrusCloudWatch.NewHook(
		viper.GetString("CLOUDWATCH_GROUP_NAME"),
		viper.GetString("CLOUDWATCH_STREAM_NAME"),
		sess,
	)
	if err != nil {
		return nil, errors.Errorf("failed to initialize CloudWatch logs: %v", err)
	}

	logger := logrus.New()
	logger.Hooks.Add(hook)
	logger.Out = os.Stdout
	logger.Formatter = logrusCloudWatch.NewProdFormatter()

	return logger, nil
}

// SetUpLogging initializes a Logrus logger and returns it.
// If the API is in a Production or Stage environment, the logger will be configured for use with AWS CloudWatch.
func SetupLogging() (*logrus.Logger, error) {
	switch GetEnvironment() {
	case config.PRODUCTION:
		fallthrough
	case config.STAGE:
		return setupCloudWatchLogger()
	case config.DEVELOPMENT:
		logger := logrus.StandardLogger()
		logger.SetLevel(logrus.DebugLevel)
		return logger, nil
	default:
		return logrus.StandardLogger(), nil
	}
}