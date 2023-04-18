package ws

import (
	awshelpers "backend/pkg/aws-helpers"
	"backend/pkg/aws-helpers/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
)

// ChatMessageRequest Received from client
type ChatMessageRequest struct {
	Text string `json:"text"`
}

func DecodeMessage[Output ChatMessageRequest](message *Message, req *Output) error {
	if fmt.Sprintf("%s/%s", message.Service, message.Action) == "chat/send" {
		return json.Unmarshal([]byte(message.Data), req)
	}
	return nil
}

// ChatMessageResponse Sending to client
type ChatMessageResponse struct {
	Text         string `json:"text"`
	ConnectionId string `json:"connectionId"`
}

func GetChatMessageResponse(connectionId string, text string) ([]byte, error) {
	data, _ := json.Marshal(ChatMessageResponse{Text: text, ConnectionId: connectionId})
	message := Message{Service: "chat", Action: "receive", Data: string(data)}
	return json.Marshal(message)
}

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
