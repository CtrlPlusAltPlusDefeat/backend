package services

import (
	awshelpers "backend/pkg/aws-helpers"
	"backend/pkg/aws-helpers/db"
	sockethelpers "backend/pkg/socket-helpers"
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

func DecodeMessage[Output ChatMessageRequest](message *sockethelpers.Message, req *Output) error {
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
	message := sockethelpers.Message{Service: "chat", Action: "receive", Data: string(data)}
	return json.Marshal(message)
}

func HandleChat(requestContext *events.APIGatewayWebsocketProxyRequestContext, message sockethelpers.Message) {
	chatMessageRequest := ChatMessageRequest{}
	err := DecodeMessage(&message, &chatMessageRequest)
	if err != nil {
		fmt.Println("Error decoding message", err)
		return
	}
	err = sendChat(requestContext.ConnectionID, chatMessageRequest)
	if err != nil {
		fmt.Println("Error when attempting to send chat", err)
	}
}

func sendChat(connectionId string, message ChatMessageRequest) error {
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	connectionTable := db.ConnectionTable{DynamoDbClient: dbClient}
	connections, err := connectionTable.GetAll()
	if err != nil {
		return err
	}
	fmt.Println("Sending ", message.Text, " to all connections")
	response, _ := GetChatMessageResponse(connectionId, message.Text)
	for _, con := range connections {
		if con.ConnectionId == connectionId {
			continue
		}
		err = sockethelpers.Send(context.TODO(), con.ConnectionId, response)
		if err != nil {
			//we got an error when sending to a client
			log.Printf("Error sending: %s", err)
		}
	}
	return nil
}
