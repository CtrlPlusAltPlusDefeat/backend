package db

import (
	"backend/pkg/models/game/settings"
	"backend/pkg/models/lobby"
	"context"
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

func (l *lobbydb) Get(lobbyId *string) (lobby.Lobby, error) {
	var result lobby.Lobby

	item, err := DynamoDb.GetItem(context.TODO(), &dynamodb.GetItemInput{TableName: aws.String(l.table),
		Key: map[string]types.AttributeValue{
			"LobbyId": &types.AttributeValueMemberS{Value: *lobbyId},
		},
	})

	if err != nil {
		log.Printf("Couldn't query %s table. Here's why: %v\n", l.table, err)
		return result, err
	}

	if len(item.Item) == 0 {
		return result, fmt.Errorf("lobby not found")
	}

	err = attributevalue.UnmarshalMap(item.Item, &result)
	if err != nil {
		log.Printf("Error unmarshalling lobby.Lobby: %s", err)
		return result, err
	}

	return result, nil
}

func (l *lobbydb) Update(lobby lobby.Lobby) error {
	_, err := DynamoDb.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(l.table),
		Key: map[string]types.AttributeValue{
			"LobbyId": &types.AttributeValueMemberS{Value: lobby.LobbyId},
		},
		ExpressionAttributeNames: map[string]string{
			"#Settings": "Settings",
			"#InGame":   "InGame",
			"#GameId":   "GameId",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":Settings": &types.AttributeValueMemberS{Value: string(lobby.Settings)},
			":InGame":   &types.AttributeValueMemberBOOL{Value: lobby.InGame},
			":GameId":   &types.AttributeValueMemberS{Value: lobby.GameId},
		},
		UpdateExpression: aws.String("set #Settings=:Settings, #InGame=:InGame, #GameId=:GameId"),
	})

	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", lobby.LobbyId, l.table, err)
	}

	return err
}

func (l *lobbydb) Add(lobbyId *string) error {

	lobbySettings, err := settings.GetDefaultWordGuess().Encode()
	if err != nil {
		return err
	}

	_, err = DynamoDb.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(l.table),
		Item: map[string]types.AttributeValue{
			"LobbyId":  &types.AttributeValueMemberS{Value: *lobbyId},
			"Settings": &types.AttributeValueMemberS{Value: string(lobbySettings)},
			"InGame":   &types.AttributeValueMemberBOOL{Value: false},
			"GameId":   &types.AttributeValueMemberS{Value: ""},
		},
	})

	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", lobbyId, l.table, err)
	}

	return err
}
