package service

import (
	awshelpers "backend/pkg/aws-helpers"
	"backend/pkg/aws-helpers/db"
	"backend/pkg/ws"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

func HandleChat(requestContext *events.APIGatewayWebsocketProxyRequestContext, message ws.Message) {
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	connectionTable := db.ConnectionTable{DynamoDbClient: dbClient}
	connections := connectionTable.GetAll()
	fmt.Println("Sending ", message.Data, " to all connections")
	for _, con := range connections {
		if con.ConnectionId == requestContext.ConnectionID {
			continue
		}
		err := ws.Send(context.TODO(), con.ConnectionId, message.Data)
		if err != nil {
			log.Fatalf("Error sending: %s", err)
		}
	}
}
