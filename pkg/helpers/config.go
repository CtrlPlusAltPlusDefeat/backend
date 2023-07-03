package helpers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"os"
)

func GetConfig() aws.Config {
	dbUrl := os.Getenv("DYNAMO_DB_URL")
	if len(dbUrl) > 0 {
		return getLocalConfig()
	}

	return getProductionConfig()
}

func getLocalConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				dbUrl := os.Getenv("DYNAMO_DB_URL")
				var endpoint string
				if len(dbUrl) > 0 {
					endpoint = dbUrl
				} else {
					endpoint = "http://dynamo-local:8000"
				}
				return aws.Endpoint{URL: endpoint}, nil

			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)

	if err != nil {
		//we panic here because this is a fatal error, we cannot continue from this
		panic(err)
	}

	return cfg
}

func getProductionConfig() aws.Config {
	defaultConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-2"))
	if err != nil {
		panic(err)
	}

	svc := secretsmanager.NewFromConfig(defaultConfig)

	keyResult, err := svc.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("BackendAccessKey"),
	})
	if err != nil {
		panic(err)
	}

	secretResult, err := svc.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("BackendSecretAccessKey"),
	})
	if err != nil {
		panic(err)
	}

	var key string = *keyResult.SecretString
	var secret string = *secretResult.SecretString

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-2"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
	)
	if err != nil {
		panic(err)
	}

	return cfg
}
