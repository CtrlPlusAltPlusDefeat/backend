package route

import (
	"backend/pkg/models"
	"backend/pkg/services"
	"log"
)

func chatHandle(socketData *SocketData) {
	chatMessageRequest := models.ChatMessageRequest{}
	err := chatMessageRequest.Decode(&socketData.message)
	if err != nil {
		log.Println("Error decoding message", err)
		return
	}
	err = services.Chat.BroadcastMessage(socketData.requestContext.ConnectionID, chatMessageRequest)
	if err != nil {
		log.Println("Error when attempting to send chat", err)
	}
}
