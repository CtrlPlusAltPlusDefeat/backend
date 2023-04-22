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

type connection struct {
}

type connectionDb struct {
	dynamo *dynamodb.Client
}

type connectionKey struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
}

type ConnectionUpdate struct {
	SessionId string `dynamodbav:"SessionId"`
}

type ConnectionItem struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
	SessionId    string `dynamodbav:"SessionId"`
}

const table = "Connection"

var Connection connection

func (c connection) GetClient() connectionDb {
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	return connectionDb{dynamo: dbClient}
}

// Add adds a connectionId to the DynamoDB table
func (conn connectionDb) Add(connection ConnectionItem) error {
	item, err := attributevalue.MarshalMap(connection)
	if err == nil {
		_, err = conn.dynamo.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String(table), Item: item,
		})
	}
	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", connection.ConnectionId, table, err)
	}
	return err
}

func (conn connectionDb) Remove(connection ConnectionItem) error {
	item, err := attributevalue.MarshalMap(connection)
	if err == nil {
		_, err = conn.dynamo.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
			TableName: aws.String(table), Key: item,
		})
	}
	if err != nil {
		log.Printf("Couldn't delete %v from the table %s. Here's why: %v\n", connection.ConnectionId, table, err)
	}
	return err
}

func (conn connectionDb) GetAll() ([]ConnectionItem, error) {

	fmt.Println("Getting all connections")

	var connections []ConnectionItem
	scan, err := conn.dynamo.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(table),
	})
	if err != nil {
		log.Printf("Error scanning db: %s", err)
		return connections, err
	}
	for _, item := range scan.Items {
		var connection ConnectionItem
		err = attributevalue.UnmarshalMap(item, &connection)
		if err != nil {
			log.Fatalf("Error unmarshalling dyanmodb map: %s", err)
		}
		connections = append(connections, connection)
	}
	return connections, nil
}

func (conn connectionDb) Get(connectionId string) (ConnectionItem, error) {
	key, err := attributevalue.MarshalMap(connectionKey{connectionId})
	var connectionItem ConnectionItem

	if err != nil {
		log.Fatalf("Error marshalling dyanmodb map: %s", err)
		return connectionItem, err
	}

	res, err := conn.dynamo.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table), Key: key,
	})
	if err != nil {
		log.Printf("Couldn't get %v from the table %s. Here's why: %v\n", connectionId, table, err)
		return connectionItem, err
	}
	err = attributevalue.UnmarshalMap(res.Item, &connectionItem)
	if err != nil {
		log.Fatalf("Error unmarshalling dyanmodb map: %s", err)
	}
	return connectionItem, err
}

func (conn connectionDb) Update(connectionId string, update ConnectionUpdate) error {
	key, err := attributevalue.MarshalMap(connectionKey{connectionId})
	if err != nil {
		log.Fatalf("Error marshalling dyanmodb map: %s", err)
		return err
	}

	updateValues, err := attributevalue.MarshalMap(update)
	if err != nil {
		log.Fatalf("Error marshalling dyanmodb map: %s", err)
		return err
	}

	_, err = conn.dynamo.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(table), Key: key, ExpressionAttributeValues: updateValues, UpdateExpression: aws.String("set info.rating = :r")})

	if err != nil {
		log.Printf("Couldn't update %v from the table %s. Here's why: %v\n", connectionId, table, err)
		return err
	}
	return err
}
