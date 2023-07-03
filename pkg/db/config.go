package db

import (
	"backend/pkg/helpers"
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
