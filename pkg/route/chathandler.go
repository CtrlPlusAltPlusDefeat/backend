package route

import (
	"backend/pkg/models"
	"backend/pkg/services"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func HandleChat(requestContext *events.APIGatewayWebsocketProxyRequestContext, message models.Wrapper) {
	chatMessageRequest := models.ChatMessageRequest{}
	err := models.DecodeMessage(&message, &chatMessageRequest)
	if err != nil {
		fmt.Println("Error decoding message", err)
		return
	}
	err = services.BroadcastMessage(requestContext.ConnectionID, chatMessageRequest)
	if err != nil {
		fmt.Println("Error when attempting to send chat", err)
	}
}
