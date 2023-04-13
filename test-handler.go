package main

import (
	"backend/aws"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

var dbClient *dynamodb.Client

type LambdaResponse struct {
	Body string `json:"body"`
}

func ListTables() ([]string, error) {
	var tableNames []string
	tables, err := dbClient.ListTables(
		context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		log.Printf("Couldn't list tables. Here's why: %v\n", err)
	} else {
		tableNames = tables.TableNames
		for index, str := range tableNames {
			fmt.Println(index, str)
		}
	}
	return tableNames, err
}
func HandleLambdaEvent() (LambdaResponse, error) {
	fmt.Println("HandleLambdaEvent")

	_, tables := ListTables()
	fmt.Println(tables)
	return LambdaResponse{Body: fmt.Sprintf("%s", tables)}, nil
}

// init is called when lambda is booted up
func init() {
	dbClient = dynamodb.NewFromConfig(aws.GetConfig())
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
