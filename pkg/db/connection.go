package db

import (
	awshelpers "backend/pkg/aws-helpers"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

type connection struct {
	dynamo *dynamodb.Client
}

type connectionKey struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
}

type ConnectionUpdate struct {
	SessionId string `dynamodbav:":SessionId"`
}

type ConnectionItem struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
	SessionId    string `dynamodbav:"SessionId"`
}

const table = "Connection"

var Connection connection

func (conn connection) GetClient() connection {
	if conn.dynamo != nil {
		return conn
	}
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	conn.dynamo = dbClient
	return conn
}

// Add adds a connectionId to the DynamoDB table
func (conn connection) Add(connectionId string) error {
	_, err := conn.dynamo.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(table), Item: map[string]types.AttributeValue{
			"ConnectionId": &types.AttributeValueMemberS{Value: connectionId},
		}})
	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", connectionId, table, err)
	}
	return err
}

func (conn connection) Remove(connectionId *string) error {
	_, err := conn.dynamo.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(table), Key: map[string]types.AttributeValue{
			"ConnectionId": &types.AttributeValueMemberS{Value: *connectionId},
		}})
	if err != nil {
		log.Printf("Couldn't delete %v from the table %s. Here's why: %v\n", connectionId, table, err)
	}
	return err
}

func (conn connection) GetAll() ([]ConnectionItem, error) {

	log.Println("Getting all connections")

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
			log.Printf("Error unmarshalling dyanmodb map: %s", err)
		}
		connections = append(connections, connection)
	}
	return connections, nil
}

func (conn connection) Get(connectionId string) (ConnectionItem, error) {
	key, err := attributevalue.MarshalMap(connectionKey{connectionId})
	var connectionItem ConnectionItem

	if err != nil {
		log.Printf("Error marshalling dyanmodb map: %s", err)
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
		log.Printf("Error unmarshalling dyanmodb map: %s", err)
	}
	return connectionItem, err
}

func (conn connection) GetBySessionId(sessionId string) ([]ConnectionItem, error) {
	var connections []ConnectionItem

	output, err := conn.dynamo.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(table),
		IndexName:              aws.String("SessionIdIndex"),
		KeyConditionExpression: aws.String("#sessionId = :v_sessionId"),
		ExpressionAttributeNames: map[string]string{
			"#sessionId": "SessionId",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v_sessionId": &types.AttributeValueMemberS{Value: sessionId},
		}})

	for _, item := range output.Items {
		var connection ConnectionItem
		err = attributevalue.UnmarshalMap(item, &connection)
		if err != nil {
			log.Printf("Error unmarshalling dyanmodb map: %s", err)
		}
		connections = append(connections, connection)
	}

	if err != nil {
		log.Printf("Couldn't Query table %s. Here's why: %v\n", table, err)
	}
	return connections, err
}

func (conn connection) Update(connectionId string, update ConnectionUpdate) error {
	log.Printf("Update %s", connectionId)
	key, err := attributevalue.MarshalMap(connectionKey{connectionId})
	if err != nil {
		log.Printf("Error marshalling dyanmodb map: %s", err)
		return err
	}

	updateValues, err := attributevalue.MarshalMap(update)
	if err != nil {
		log.Printf("Error marshalling dyanmodb map: %s", err)
		return err
	}

	_, err = conn.dynamo.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table),
		Key:                       key,
		ExpressionAttributeValues: updateValues,
		UpdateExpression:          aws.String("set SessionId = :SessionId"),
		ReturnValues:              types.ReturnValueUpdatedNew,
	})

	if err != nil {
		log.Printf("Couldn't update %v from the table %s. Here's why: %v\n", connectionId, table, err)
	}
	return err
}
