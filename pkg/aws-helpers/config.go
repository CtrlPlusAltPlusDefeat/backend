package aws_helpers

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-secretsmanager-caching-go/secretcache"
	"os"
	"log"
)

var (
	 secretCache, _ = secretcache.New()
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
  
	key, _ := secretCache.GetSecretString("BackendAccessKey")
  
	log.Printf("GetConfig #6")
  
	secret, _ := secretCache.GetSecretString("BackendSecretAccessKey")

	log.Printf("GetConfig #7")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
	)

	log.Printf("GetConfig #8")

	if err != nil {
		log.Printf("GetConfig #9")
    
		//we panic here because this is a fatal error, we cannot continue from this
		panic(err)
	}
	
	log.Printf("GetConfig #10")
  
	return cfg
}
