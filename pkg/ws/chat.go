package ws

import (
	awshelpers "backend/pkg/aws-helpers"
	"backend/pkg/aws-helpers/db"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

func HandleChat(requestContext *events.APIGatewayWebsocketProxyRequestContext, message Message) {
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	connectionTable := db.ConnectionTable{DynamoDbClient: dbClient}
	connections := connectionTable.GetAll()

	chatMessageRequest := ChatMessageRequest{}
	err := DecodeMessage(&message, &chatMessageRequest)
	if err != nil {
		fmt.Println("Error decoding message", err)
		return
	}

	fmt.Println("Sending ", chatMessageRequest.Text, " to all connections")
	response, _ := GetChatMessageResponse(requestContext.ConnectionID, chatMessageRequest.Text)
	for _, con := range connections {
		if con.ConnectionId == requestContext.ConnectionID {
			continue
		}
		err := Send(context.TODO(), con.ConnectionId, response)
		if err != nil {
			log.Fatalf("Error sending: %s", err)
		}
	}
}
