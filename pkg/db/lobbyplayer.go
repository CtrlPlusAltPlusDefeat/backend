package db

import (
	awshelpers "backend/pkg/aws-helpers"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"log"
)

type lobbyplayer struct {
	dynamo *dynamodb.Client
	table  string
}

type Player struct {
	LobbyId      string `dynamodbav:"LobbyId"`
	SessionId    string `dynamodbav:"SessionId"`
	ConnectionId string `dynamodbav:"ConnectionId"`
	Id           string `dynamodbav:"Id"`
	Name         string `dynamodbav:"Name"`
	Points       int32  `dynamodbav:"Points"`
	IsAdmin      bool   `dynamodbav:"IsAdmin"`
}

var LobbyPlayer = lobbyplayer{dynamo: nil, table: "LobbyPlayer"}

func (l *lobbyplayer) getClient() {
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	l.dynamo = dbClient

}

func (l *lobbyplayer) Add(lobbyId *string, sessionId *string, connectionId *string, isAdmin bool) (Player, error) {
	var player Player
	if l.dynamo == nil {
		l.getClient()
	}
	item, err := l.dynamo.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(l.table), Item: map[string]types.AttributeValue{
			"LobbyId":      &types.AttributeValueMemberS{Value: *lobbyId},
			"SessionId":    &types.AttributeValueMemberS{Value: *sessionId},
			"ConnectionId": &types.AttributeValueMemberS{Value: *connectionId},
			"Id":           &types.AttributeValueMemberS{Value: uuid.New().String()},
			"Name":         &types.AttributeValueMemberS{Value: ""},
			"Points":       &types.AttributeValueMemberN{Value: "0"},
			"IsAdmin":      &types.AttributeValueMemberBOOL{Value: isAdmin},
		}})
	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", lobbyId, l.table, err)
	}

	err = attributevalue.UnmarshalMap(item.Attributes, &player)

	if err != nil {
		log.Printf("Error unmarshalling dyanmodb map: %s", err)
	}
	return player, err
}

func (l *lobbyplayer) Remove(lobbyId *string, sessionId *string) error {
	if l.dynamo == nil {
		l.getClient()
	}
	_, err := l.dynamo.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(l.table), Key: map[string]types.AttributeValue{
			"LobbyId":   &types.AttributeValueMemberS{Value: *lobbyId},
			"SessionId": &types.AttributeValueMemberS{Value: *sessionId},
		}})
	if err != nil {
		log.Printf("Couldn't delete %s from %s table. Here's why: %v\n", lobbyId, l.table, err)
	}
	return err
}

func (l *lobbyplayer) GetPlayers(lobbyId *string) ([]Player, error) {
	var players []Player
	if l.dynamo == nil {
		l.getClient()
	}
	query, err := l.dynamo.Query(context.TODO(), &dynamodb.QueryInput{TableName: aws.String(l.table),
		KeyConditionExpression: aws.String("LobbyId=:LobbyId"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":LobbyId": &types.AttributeValueMemberS{Value: *lobbyId},
		},
	})
	if err != nil {
		log.Printf("Couldn't query %s table. Here's why: %v\n", l.table, err)
		return players, err
	}
	for _, item := range query.Items {
		var player Player
		err = attributevalue.UnmarshalMap(item, &player)
		if err != nil {
			log.Printf("Error unmarshalling dyanmodb map: %s", err)
		}
		players = append(players, player)
	}
	return players, nil
}

func (l *lobbyplayer) UpdateName(lobbyId *string, sessionId *string, name *string) (Player, error) {
	var player Player
	if l.dynamo == nil {
		l.getClient()
	}
	item, err := l.dynamo.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(l.table),
		Key: map[string]types.AttributeValue{
			"SessionId": &types.AttributeValueMemberS{Value: *sessionId},
			"LobbyId":   &types.AttributeValueMemberS{Value: *lobbyId},
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":Name": &types.AttributeValueMemberS{Value: *name},
		},
		ExpressionAttributeNames: map[string]string{
			"#Name": "Name",
		},
		UpdateExpression: aws.String("set #Name=:Name"),
		ReturnValues:     types.ReturnValueAllNew,
	})
	if err != nil {
		log.Printf("Error updating sessionId %s to name: %s. %s", *sessionId, *name, err)

		return player, err
	}
	err = attributevalue.UnmarshalMap(item.Attributes, &player)
	if err != nil {
		log.Printf("Error unmarshalling dyanmodb map: %s", err)
	}
	return player, err
}
