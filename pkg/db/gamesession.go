package db

import (
	"backend/pkg/models/game"
	"backend/pkg/models/lobby"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"strconv"
)

type gamedb struct {
	table string
}

var GameSession = gamedb{table: "GameSession"}

func (g *gamedb) Get(lobbyId *string) (lobby.Lobby, error) {
	var result lobby.Lobby

	item, err := DynamoDb.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(g.table),
		Key: map[string]types.AttributeValue{
			"LobbyId": &types.AttributeValueMemberS{Value: *lobbyId},
		},
	})

	if err != nil {
		log.Printf("Couldn't query %s table. Here's why: %v\n", g.table, err)
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

func (g *gamedb) Add(lobbyId *string, gameSessionId *string, gameTypeId game.Id) (*game.Session, error) {
	result := game.Session{
		LobbyId: *lobbyId, GameSessionId: *gameSessionId, GameTypeId: gameTypeId,
	}
	_, err := DynamoDb.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(g.table),
		Item: map[string]types.AttributeValue{
			"LobbyId":       &types.AttributeValueMemberS{Value: *lobbyId},
			"GameSessionId": &types.AttributeValueMemberS{Value: *gameSessionId},
			"GameTypeId":    &types.AttributeValueMemberN{Value: strconv.Itoa(int(gameTypeId))},
		},
	})

	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", *lobbyId, g.table, err)
		return nil, err
	}

	return &result, nil

}
