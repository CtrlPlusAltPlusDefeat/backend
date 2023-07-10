package db

import (
	"backend/pkg/game/wordguess"
	"backend/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
)

type lobbydb struct {
	table string
}

var Lobby = lobbydb{table: "Lobby"}

type lobbyItem struct {
	LobbyId       string `dynamodbav:"LobbyId"`
	Settings      string `dynamodbav:"Settings"`
	InGame        bool   `dynamodbav:"InGame"`
	GameSessionId string `dynamodbav:"GameSessionId"`
}

func (l *lobbydb) Get(lobbyId *string) (models.Lobby, error) {
	var lobby models.Lobby
	var res lobbyItem

	item, err := DynamoDb.GetItem(context.TODO(), &dynamodb.GetItemInput{TableName: aws.String(l.table),
		Key: map[string]types.AttributeValue{
			"LobbyId": &types.AttributeValueMemberS{Value: *lobbyId},
		},
	})

	if err != nil {
		log.Printf("Couldn't query %s table. Here's why: %v\n", l.table, err)
		return lobby, err
	}

	if len(item.Item) == 0 {
		return lobby, fmt.Errorf("lobby not found")
	}

	err = attributevalue.UnmarshalMap(item.Item, &res)
	if err != nil {
		log.Printf("Error unmarshalling lobby.Lobby: %s", err)
		return lobby, err
	}

	var settings models.Settings
	err = json.Unmarshal([]byte(res.Settings), &settings)
	lobby = models.Lobby{LobbyId: res.LobbyId, Settings: settings, InGame: res.InGame, GameSessionId: res.GameSessionId}

	return lobby, nil
}

func (l *lobbydb) Update(lobby *models.Lobby) error {

	settings, err := lobby.Settings.Encode()
	if err != nil {
		return err
	}
	_, err = DynamoDb.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(l.table),
		Key: map[string]types.AttributeValue{
			"LobbyId": &types.AttributeValueMemberS{Value: lobby.LobbyId},
		},
		ExpressionAttributeNames: map[string]string{
			"#Settings":      "Settings",
			"#InGame":        "InGame",
			"#GameSessionId": "GameSessionId",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":Settings":      &types.AttributeValueMemberS{Value: string(settings)},
			":InGame":        &types.AttributeValueMemberBOOL{Value: lobby.InGame},
			":GameSessionId": &types.AttributeValueMemberS{Value: lobby.GameSessionId},
		},
		UpdateExpression: aws.String("set #Settings=:Settings, #InGame=:InGame, #GameSessionId=:GameSessionId"),
	})

	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", lobby.LobbyId, l.table, err)
	}

	return err
}

func (l *lobbydb) Add(lobbyId *string) error {

	lobbySettings, err := wordguess.GetDefaultSettings().Encode()
	if err != nil {
		return err
	}

	_, err = DynamoDb.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(l.table),
		Item: map[string]types.AttributeValue{
			"LobbyId":       &types.AttributeValueMemberS{Value: *lobbyId},
			"Settings":      &types.AttributeValueMemberS{Value: string(lobbySettings)},
			"InGame":        &types.AttributeValueMemberBOOL{Value: false},
			"GameSessionId": &types.AttributeValueMemberS{Value: ""},
		},
	})

	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", lobbyId, l.table, err)
	}

	return err
}
