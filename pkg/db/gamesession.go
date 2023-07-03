package db

import (
	"backend/pkg/models/game"
	"context"
	"encoding/json"
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
	var state game.SessionState
	var gInfo game.SessionInfo

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

	teamsJson := game.EncodedTeamArray(item.Item["Teams"].(*types.AttributeValueMemberS).Value)
	gState := game.EncodedTeamArray(item.Item["Game"].(*types.AttributeValueMemberS).Value)
	err = attributevalue.UnmarshalMap(item.Item, &state)
	if err != nil {
		log.Printf("Error unmarshalling state.GameState: %s", err)
		return nil, err
	}
	err = attributevalue.UnmarshalMap(item.Item, &gInfo)
	if err != nil {
		log.Printf("Error unmarshalling game.Session: %s", err)
		return nil, err
	}

	return &game.Session{
		State: &state,
		Info:  &gInfo,
		Teams: *teamsJson.Decode(),
		Game:  json.RawMessage(gState),
	}, nil
}

func (g *gamedb) Add(sessions *game.Session) (*game.Session, error) {
	encoded, err := sessions.Teams.Encode()
	if err != nil {
		log.Printf("Couldn't encode teams for %s. Here's why: %v\n", sessions.Info.LobbyId, err)
		return sessions, err
	}
	_, err = DynamoDb.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(g.table),
		Item: map[string]types.AttributeValue{
			"LobbyId":       &types.AttributeValueMemberS{Value: sessions.Info.LobbyId},
			"GameSessionId": &types.AttributeValueMemberS{Value: sessions.Info.GameSessionId},
			"GameTypeId":    &types.AttributeValueMemberN{Value: strconv.Itoa(int(sessions.Info.GameTypeId))},
			"State":         &types.AttributeValueMemberS{Value: string(sessions.State.State)},
			"CurrentTurn":   &types.AttributeValueMemberS{Value: string(sessions.State.CurrentTurn)},
			"Teams":         &types.AttributeValueMemberS{Value: string(*encoded)},
			"Game":          &types.AttributeValueMemberS{Value: string(sessions.Game)},
		},
	})

	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", sessions.Info.LobbyId, g.table, err)
		return nil, err
	}

	return sessions, nil

}
