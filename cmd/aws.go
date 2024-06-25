package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-frame/config"
	"github.com/go-frame/internals/lib/logger"
)

func setupAWS(conf *config.Config, appLogger logger.Logger) *session.Session {
	accessKey := conf.Aws.AwsAccessKeyId
	accessSecret := conf.Aws.AwsSecretAccessKey
	defaultRegion := conf.Aws.AwsDefaultRegion
	bucketName := conf.Aws.AwsStorageBucketName
	var awsConfig *aws.Config
	if accessKey == "" || accessSecret == "" || defaultRegion == "" || bucketName == "" {
		appLogger.Info("aws configuration missing")
	} else {
		awsConfig = &aws.Config{
			Region:      aws.String(defaultRegion),
			Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
		}
	}

	return session.Must(session.NewSession(awsConfig))
}
