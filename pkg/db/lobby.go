package db

import (
	awshelpers "backend/pkg/aws-helpers"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

type lobby struct {
	dynamo *dynamodb.Client
	table  string
}

var Lobby = lobby{dynamo: nil, table: "Lobby"}

func (l *lobby) getClient() {
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	l.dynamo = dbClient

}

func (l *lobby) Add(lobbyId string) error {
	if l.dynamo == nil {
		l.getClient()
	}
	_, err := l.dynamo.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(l.table), Item: map[string]types.AttributeValue{
			"LobbyId": &types.AttributeValueMemberS{Value: lobbyId},
		}})
	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", lobbyId, l.table, err)
	}
	return err
}
