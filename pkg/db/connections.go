package db

import (
	awshelpers "backend/pkg/aws-helpers"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

type ConnectionDb struct {
	DynamoDbClient *dynamodb.Client
}

type Connection struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
}

const table = "Connection"

func GetConnectionDb() ConnectionDb {
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	return ConnectionDb{DynamoDbClient: dbClient}
}

// Add adds a connectionId to the DynamoDB table
func (conn ConnectionDb) Add(connection Connection) error {
	item, err := attributevalue.MarshalMap(connection)
	if err == nil {
		_, err = conn.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String(table), Item: item,
		})
	}
	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", connection.ConnectionId, table, err)
	}
	return err
}

func (conn ConnectionDb) Remove(connection Connection) error {
	item, err := attributevalue.MarshalMap(connection)
	if err == nil {
		_, err = conn.DynamoDbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
			TableName: aws.String(table), Key: item,
		})
	}
	if err != nil {
		log.Printf("Couldn't delete %v from the table %s. Here's why: %v\n", connection.ConnectionId, table, err)
	}
	return err
}

func (conn ConnectionDb) GetAll() ([]Connection, error) {
	fmt.Println("Getting all connections")

	var connections []Connection
	scan, err := conn.DynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(table),
	})
	if err != nil {
		log.Printf("Error scanning db: %s", err)
		return connections, err
	}
	for _, item := range scan.Items {
		var connection Connection
		err = attributevalue.UnmarshalMap(item, &connection)
		if err != nil {
			log.Fatalf("Error unmarshalling dyanmodb map: %s", err)
		}
		connections = append(connections, connection)
	}
	return connections, nil
}
