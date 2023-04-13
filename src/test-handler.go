package main

import (
	"backend/src/aws"
	dynamodb2 "backend/src/aws/dynamodb"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var dbClient *dynamodb.Client
var tableBasics dynamodb2.TableBasics

type LambdaResponse struct {
	Body string `json:"body"`
}

type TableList struct {
	Tables []string `json:"tables"`
}

func HandleLambdaEvent() (LambdaResponse, error) {
	fmt.Println("Entered HandleLambdaEvent")

	tables, err := tableBasics.ListTables()
	if err != nil {
		return LambdaResponse{}, err
	}

	tableList := TableList{Tables: tables}
	encodedJson, _ := json.Marshal(tableList)
	return LambdaResponse{Body: fmt.Sprintf("%s", encodedJson)}, nil
}

// init is called when lambda is booted up
func init() {
	dbClient = dynamodb.NewFromConfig(aws.GetConfig())
	tableBasics = dynamodb2.TableBasics{DynamoDbClient: dbClient}
	_, err := tableBasics.CreateTable("TestTable-Code")
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
