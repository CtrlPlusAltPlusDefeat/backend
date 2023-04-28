package main

import (
	awshelpers "backend/pkg/aws-helpers"
	"backend/pkg/db"
	"backend/pkg/route"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

/*
*
init is called when lambda starts up,
spin up a dynamodb client and inject into db package
*/
func init() {
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	db.DynamoDb = dbClient
}

func main() {
	lambda.Start(route.ConnectHandler)
}
