package db

import (
	"backend/pkg/models"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

type connection struct {
	table string
}

var Connection = connection{"Connection"}

// Add adds a connectionId to the DynamoDB table
func (conn connection) Add(connectionId string) error {
	_, err := DynamoDb.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(conn.table), Item: map[string]types.AttributeValue{
			"ConnectionId": &types.AttributeValueMemberS{Value: connectionId},
		}})
	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", connectionId, conn.table, err)
	}
	return err
}

func (conn connection) Remove(connectionId *string) error {
	_, err := DynamoDb.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(conn.table), Key: map[string]types.AttributeValue{
			"ConnectionId": &types.AttributeValueMemberS{Value: *connectionId},
		}})
	if err != nil {
		log.Printf("Couldn't delete %v from the table %s. Here's why: %v\n", connectionId, conn.table, err)
	}
	return err
}

func (conn connection) GetAll() ([]models.Connection, error) {

	log.Println("Getting all connections")

	var connections []models.Connection
	scan, err := DynamoDb.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(conn.table),
	})
	if err != nil {
		log.Printf("Error scanning db: %s", err)
		return connections, err
	}
	for _, item := range scan.Items {
		var c models.Connection
		err = attributevalue.UnmarshalMap(item, &c)
		if err != nil {
			log.Printf("Error unmarshalling dyanmodb map: %s", err)
		}
		connections = append(connections, c)
	}
	return connections, nil
}

func (conn connection) Get(connectionId string) (models.Connection, error) {
	var connectionItem models.Connection
	res, err := DynamoDb.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(conn.table), Key: map[string]types.AttributeValue{
			"ConnectionId": &types.AttributeValueMemberS{Value: connectionId},
		},
	})
	if err != nil {
		log.Printf("Couldn't get %v from the table %s. Here's why: %v\n", connectionId, conn.table, err)
		return connectionItem, err
	}
	err = attributevalue.UnmarshalMap(res.Item, &connectionItem)
	if err != nil {
		log.Printf("Error unmarshalling dyanmodb map: %s", err)
	}
	return connectionItem, err
}

func (conn connection) GetBySessionId(sessionId string) ([]models.Connection, error) {
	var connections []models.Connection

	output, err := DynamoDb.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(conn.table),
		IndexName:              aws.String("SessionIdIndex"),
		KeyConditionExpression: aws.String("#sessionId = :v_sessionId"),
		ExpressionAttributeNames: map[string]string{
			"#sessionId": "SessionId",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v_sessionId": &types.AttributeValueMemberS{Value: sessionId},
		}})

	if err != nil {
		log.Printf("Couldn't get %v from the table %s. Here's why: %v\n", sessionId, conn.table, err)
		return connections, err
	}

	for _, item := range output.Items {
		var connection models.Connection
		err = attributevalue.UnmarshalMap(item, &connection)
		if err != nil {
			log.Printf("Error unmarshalling dyanmodb map: %s", err)
		}
		connections = append(connections, connection)
	}

	if err != nil {
		log.Printf("Couldn't Query table %s. Here's why: %v\n", conn.table, err)
	}
	return connections, err
}

func (conn connection) Update(connectionId string, sessionId string) error {

	_, err := DynamoDb.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(conn.table),
		Key: map[string]types.AttributeValue{
			"ConnectionId": &types.AttributeValueMemberS{Value: connectionId},
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":SessionId": &types.AttributeValueMemberS{Value: sessionId},
		},
		UpdateExpression: aws.String("set SessionId = :SessionId"),
		ReturnValues:     types.ReturnValueUpdatedNew,
	})

	if err != nil {
		log.Printf("Couldn't update %v from the table %s. Here's why: %v\n", connectionId, conn.table, err)
	}
	return err
}
