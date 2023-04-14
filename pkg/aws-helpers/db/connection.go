package db

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

type ConnectionTable struct {
	DynamoDbClient *dynamodb.Client
}

type Connection struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
}

const table = "Connection"

// Add adds a connectionId to the DynamoDB table
func (conn ConnectionTable) Add(connection Connection) error {
	item, err := attributevalue.MarshalMap(connection)
	if err != nil {
		panic(err)
	}
	_, err = conn.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(table), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", connection.ConnectionId, table, err)
	}
	return err
}

func (conn ConnectionTable) Remove(connectionId string) error {
	_, err := conn.DynamoDbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(table), ConditionExpression: aws.String(fmt.Sprintf("ConnectionId=%s", connectionId)),
	})
	if err != nil {
		log.Printf("Couldn't delete %v from the table %s. Here's why: %v\n", connectionId, table, err)
	}
	return err
}

func (conn ConnectionTable) GetAll() []Connection {
	fmt.Println("Getting all connections")

	var connections []Connection
	scan, err := conn.DynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(table),
	})
	if err != nil {
		log.Fatalf("Error scanning db: %s", err)
	}
	for _, item := range scan.Items {
		var connection Connection
		err = attributevalue.UnmarshalMap(item, connection)
		if err != nil {
			log.Fatalf("Error unmarshalling: %s", err)
		}
		connections = append(connections, connection)
	}
	return connections
}
