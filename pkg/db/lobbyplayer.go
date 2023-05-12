package db

import (
	"backend/pkg/models/lobby"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"log"
)

type lobbyplayer struct {
	table string
}

var LobbyPlayer = lobbyplayer{table: "LobbyPlayer"}

func (l *lobbyplayer) Add(lobbyId *string, sessionId *string, connectionId *string, name string, isAdmin bool) (lobby.Player, error) {
	var player lobby.Player

	item, err := DynamoDb.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(l.table),
		Key: map[string]types.AttributeValue{
			"LobbyId":   &types.AttributeValueMemberS{Value: *lobbyId},
			"SessionId": &types.AttributeValueMemberS{Value: *sessionId},
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":ConnectionId": &types.AttributeValueMemberS{Value: *connectionId},
			":Id":           &types.AttributeValueMemberS{Value: uuid.New().String()},
			":Name":         &types.AttributeValueMemberS{Value: name},
			":Points":       &types.AttributeValueMemberN{Value: "0"},
			":IsAdmin":      &types.AttributeValueMemberBOOL{Value: isAdmin},
			":IsOnline":     &types.AttributeValueMemberBOOL{Value: true},
		},
		ExpressionAttributeNames: map[string]string{
			"#ConnectionId": "ConnectionId",
			"#IsAdmin":      "IsAdmin",
			"#IsOnline":     "IsOnline",
			"#Id":           "Id",
			"#Name":         "Name",
			"#Points":       "Points",
		},
		UpdateExpression: aws.String("set #ConnectionId=:ConnectionId, #IsAdmin=:IsAdmin, #IsOnline=:IsOnline, #Id=if_not_exists(#Id, :Id), #Name=if_not_exists(#Name, :Name), #Points=if_not_exists(#Points, :Points)"),
		ReturnValues:     types.ReturnValueAllNew,
	})

	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", lobbyId, l.table, err)
	}

	err = attributevalue.UnmarshalMap(item.Attributes, &player)

	if err != nil {
		log.Printf("Error unmarshalling dyanmodb map: %s", err)
	}
	return player, err
}

func (l *lobbyplayer) GetPlayers(lobbyId *string) ([]lobby.Player, error) {
	var players []lobby.Player

	query, err := DynamoDb.Query(context.TODO(), &dynamodb.QueryInput{TableName: aws.String(l.table),
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
		var player lobby.Player
		err = attributevalue.UnmarshalMap(item, &player)
		if err != nil {
			log.Printf("Error unmarshalling dyanmodb map: %s", err)
		}
		players = append(players, player)
	}
	return players, nil
}

func (l *lobbyplayer) Get(lobbyId *string, sessionId *string) (lobby.Player, error) {
	var player lobby.Player

	item, err := DynamoDb.GetItem(context.TODO(), &dynamodb.GetItemInput{TableName: aws.String(l.table),
		Key: map[string]types.AttributeValue{
			"SessionId": &types.AttributeValueMemberS{Value: *sessionId},
			"LobbyId":   &types.AttributeValueMemberS{Value: *lobbyId},
		},
	})
	if err != nil {
		log.Printf("Couldn't query %s table. Here's why: %v\n", l.table, err)
		return player, err
	}

	if len(item.Item) == 0 {
		return player, fmt.Errorf("player not found")
	}

	err = attributevalue.UnmarshalMap(item.Item, &player)
	if err != nil {
		log.Printf("Error unmarshalling lobby.Player: %s", err)
	}
	return player, nil
}

func (l *lobbyplayer) UpdateOnline(lobbyId *string, sessionId *string, online bool) (lobby.Player, error) {
	var player lobby.Player

	item, err := DynamoDb.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(l.table),
		Key: map[string]types.AttributeValue{
			"SessionId": &types.AttributeValueMemberS{Value: *sessionId},
			"LobbyId":   &types.AttributeValueMemberS{Value: *lobbyId},
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":IsOnline": &types.AttributeValueMemberBOOL{Value: online},
		},
		ExpressionAttributeNames: map[string]string{
			"#IsOnline": "IsOnline",
		},
		UpdateExpression: aws.String("set #IsOnline=:IsOnline"),
		ReturnValues:     types.ReturnValueAllNew,
	})

	if err != nil {
		return player, err
	}

	err = attributevalue.UnmarshalMap(item.Attributes, &player)

	if err != nil {
		log.Printf("Error unmarshalling dyanmodb map: %s", err)
	}

	return player, err
}

func (l *lobbyplayer) UpdateName(lobbyId *string, sessionId *string, name *string) (lobby.Player, error) {
	var player lobby.Player

	item, err := DynamoDb.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
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
