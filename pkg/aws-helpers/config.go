package aws_helpers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log"
	"os"
)

func GetConfig() aws.Config {
	log.Printf("GetConfig #1")

	dbUrl := os.Getenv("DYNAMO_DB_URL")

	log.Printf("GetConfig #2")

	if len(dbUrl) > 0 {
		log.Printf("GetConfig #3 - Local")
		return getLocalConfig()
	}

	log.Printf("GetConfig #3 - Prod")

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
	log.Printf("GetConfig #5")

	defaultConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-2"))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("GetConfig #6")

	svc := secretsmanager.NewFromConfig(defaultConfig)

	log.Printf("GetConfig #7")

	result, err := svc.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("BackendAccessKey"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("GetConfig #8")

	var key string = *result.SecretString

	log.Printf("GetConfig #9")

	result, err = svc.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String("BackendSecretAccessKey"),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("GetConfig #10")

	var secret string = *result.SecretString

	log.Printf("GetConfig #11")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-2"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
	)

	log.Printf("GetConfig #12")

	if err != nil {
		log.Printf("GetConfig #13")
		log.Printf(err.Error())

		panic(err)
	}

	log.Printf("GetConfig #14")

	return cfg
}
