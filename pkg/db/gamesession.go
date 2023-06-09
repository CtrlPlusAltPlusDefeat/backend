package db

import (
	"backend/pkg/models/game"
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

func (g *gamedb) Get(lobbyId *string, gameSessionId *string) (*game.Session, error) {
	var gSession game.Session
	var gState game.State

	item, err := DynamoDb.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(g.table),
		Key: map[string]types.AttributeValue{
			"LobbyId":       &types.AttributeValueMemberS{Value: *lobbyId},
			"GameSessionId": &types.AttributeValueMemberS{Value: *gameSessionId},
		},
	})

	if err != nil {
		log.Printf("Couldn't query %s table. Here's why: %v\n", g.table, err)
		return nil, err
	}

	if len(item.Item) == 0 {
		return nil, fmt.Errorf("lobby not found")
	}

	err = attributevalue.UnmarshalMap(item.Item, &gState)
	if err != nil {
		log.Printf("Error unmarshalling state.GameState: %s", err)
		return nil, err
	}
	err = attributevalue.UnmarshalMap(item.Item, &gSession)
	if err != nil {
		log.Printf("Error unmarshalling game.Session: %s", err)
		return nil, err
	}
	gSession.GameState = &gState
	return &gSession, nil
}

func (g *gamedb) Add(sessions *game.Session) (*game.Session, error) {
	encoded, err := sessions.GameState.Teams.Encode()
	if err != nil {
		log.Printf("Couldn't encode teams for %s. Here's why: %v\n", sessions.LobbyId, err)
		return sessions, err
	}
	_, err = DynamoDb.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(g.table),
		Item: map[string]types.AttributeValue{
			"LobbyId":       &types.AttributeValueMemberS{Value: sessions.LobbyId},
			"GameSessionId": &types.AttributeValueMemberS{Value: sessions.GameSessionId},
			"GameTypeId":    &types.AttributeValueMemberN{Value: strconv.Itoa(int(sessions.GameTypeId))},
			"State":         &types.AttributeValueMemberS{Value: string(sessions.GameState.State)},
			"CurrentTurn":   &types.AttributeValueMemberS{Value: string(sessions.GameState.CurrentTurn)},
			"Teams":         &types.AttributeValueMemberS{Value: string(*encoded)},
		},
	})

	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", sessions.LobbyId, g.table, err)
		return nil, err
	}

	return sessions, nil

}
