package db

import (
	"backend/pkg/models"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"strconv"
	"time"
)

type lobbychat struct {
	table string
}

var LobbyChat = lobbychat{table: "LobbyChat"}

func (l *lobbychat) Add(lobbyId *string, playerId *string, message *string) (models.Chat, error) {
	c := models.Chat{
		Message:   *message,
		Timestamp: time.Now().Unix(),
		PlayerId:  *playerId,
		LobbyId:   *lobbyId,
	}

	_, err := DynamoDb.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(l.table),
		Item: map[string]types.AttributeValue{
			"LobbyId":   &types.AttributeValueMemberS{Value: c.LobbyId},
			"PlayerId":  &types.AttributeValueMemberS{Value: c.PlayerId},
			"Timestamp": &types.AttributeValueMemberN{Value: strconv.FormatInt(c.Timestamp, 10)},
			"Message":   &types.AttributeValueMemberS{Value: c.Message},
		},
		ReturnValues: types.ReturnValueNone,
	})

	if err != nil {
		log.Printf("Couldn't add %s to %s table. Here's why: %v\n", lobbyId, l.table, err)
	}

	return c, err
}

func (l *lobbychat) Get(lobbyId *string, timestamp int64) ([]models.Chat, error) {
	var chats []models.Chat

	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}

	query, err := DynamoDb.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(l.table),
		KeyConditionExpression: aws.String("LobbyId = :LobbyId AND #Timestamp <= :Timestamp"),
		ExpressionAttributeNames: map[string]string{
			"#Timestamp": "Timestamp",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":LobbyId":   &types.AttributeValueMemberS{Value: *lobbyId},
			":Timestamp": &types.AttributeValueMemberN{Value: strconv.FormatInt(timestamp, 10)},
		},
		ScanIndexForward: aws.Bool(false),
		Limit:            aws.Int32(50),
	})

	if err != nil {
		log.Printf("Couldn't query %s table. Here's why: %v\n", l.table, err)
		return chats, err
	}

	for _, item := range query.Items {
		var c models.Chat
		err = attributevalue.UnmarshalMap(item, &c)

		if err != nil {
			log.Printf("Error unmarshalling dyanmodb map: %s", err)
		}

		chats = append(chats, c)
	}

	return chats, nil
}
