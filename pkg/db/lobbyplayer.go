package db

import (
	"backend/pkg/models"
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

func (l *lobbyplayer) Add(lobbyId *string, sessionId *string, connectionId *string, name string, isAdmin bool) (models.Player, error) {
	var p models.Player

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

	err = attributevalue.UnmarshalMap(item.Attributes, &p)

	if err != nil {
		log.Printf("Error unmarshalling dyanmodb map: %s", err)
	}
	return p, err
}

func (l *lobbyplayer) GetPlayers(lobbyId *string) ([]models.Player, error) {
	var players []models.Player

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
		var p models.Player
		err = attributevalue.UnmarshalMap(item, &p)
		if err != nil {
			log.Printf("Error unmarshalling dyanmodb map: %s", err)
		}
		players = append(players, p)
	}
	return players, nil
}

func (l *lobbyplayer) Get(lobbyId *string, sessionId *string) (models.Player, error) {
	var p models.Player

	item, err := DynamoDb.GetItem(context.TODO(), &dynamodb.GetItemInput{TableName: aws.String(l.table),
		Key: map[string]types.AttributeValue{
			"SessionId": &types.AttributeValueMemberS{Value: *sessionId},
			"LobbyId":   &types.AttributeValueMemberS{Value: *lobbyId},
		},
	})
	if err != nil {
		log.Printf("Couldn't query %s table. Here's why: %v\n", l.table, err)
		return p, err
	}

	if len(item.Item) == 0 {
		return p, fmt.Errorf("player not found")
	}

	err = attributevalue.UnmarshalMap(item.Item, &p)
	if err != nil {
		log.Printf("Error unmarshalling lobby.Player: %s", err)
	}
	return p, nil
}

func (l *lobbyplayer) UpdateOnline(lobbyId *string, sessionId *string, online bool) (models.Player, error) {
	var p models.Player

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
		return p, err
	}

	err = attributevalue.UnmarshalMap(item.Attributes, &p)

	if err != nil {
		log.Printf("Error unmarshalling dyanmodb map: %s", err)
	}

	return p, err
}

// GetLobbiesBySessionId returns all lobbies that a players sessionId is in
func (l *lobbyplayer) GetLobbiesBySessionId(sessionId *string) ([]models.Player, error) {
	var lobbyPlayers []models.Player

	output, err := DynamoDb.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(l.table),
		IndexName:              aws.String("SessionIdIndex"),
		KeyConditionExpression: aws.String("#sessionId = :v_sessionId"),
		ExpressionAttributeNames: map[string]string{
			"#sessionId": "SessionId",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v_sessionId": &types.AttributeValueMemberS{Value: *sessionId},
		}})

	if err != nil {
		log.Printf("Couldn't get %v from the table %s. Here's why: %v\n", sessionId, l.table, err)
		return lobbyPlayers, err
	}

	for _, item := range output.Items {
		var p models.Player
		err = attributevalue.UnmarshalMap(item, &p)
		if err != nil {
			log.Printf("Error unmarshalling dyanmodb map: %s", err)
		}
		lobbyPlayers = append(lobbyPlayers, p)
	}

	if err != nil {
		log.Printf("Couldn't Query table %s. Here's why: %v\n", l.table, err)
	}
	return lobbyPlayers, err
}
