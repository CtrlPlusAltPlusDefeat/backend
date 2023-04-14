package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

type connection struct {
	DynamoDbClient *dynamodb.Client
}

const table = "connections"

// Add adds a connectionId to the DynamoDB table
func (basics TableBasics) Add(connectionId string) error {
	item, err := attributevalue.MarshalMap(connectionId)
	if err != nil {
		panic(err)
	}
	_, err = basics.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(table), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", connectionId, table, err)
	}
	return err
}

func (basics TableBasics) Remove(connectionId string) error {
	_, err := basics.DynamoDbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(table), ConditionExpression: aws.String(fmt.Sprintf("ConnectionId=%s", connectionId)),
	})
	if err != nil {
		log.Printf("Couldn't delete %v from the table %s. Here's why: %v\n", connectionId, table, err)
	}
	return err

}
