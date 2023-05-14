package db

import (
	helpers "backend/pkg/aws-helpers"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	DynamoDb *dynamodb.Client
)

func Configure() {
	config := helpers.GetConfig()
	client := dynamodb.NewFromConfig(config)

	DynamoDb = client
}
