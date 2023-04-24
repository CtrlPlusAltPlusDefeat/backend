package route

import (
	"backend/pkg/models"
	"backend/pkg/services"
	"log"
)

func chatHandle(socketData *models.SocketData) {
	log.Printf("chatHandle: %s", socketData.Message.Action)

	chatMessageRequest := models.ChatMessageRequest{}
	err := chatMessageRequest.Decode(&socketData.Message)
	if err != nil {
		log.Println("Error decoding message", err)
		return
	}
	err = services.Chat.BroadcastMessage(socketData.RequestContext.ConnectionID, chatMessageRequest)
	if err != nil {
		log.Println("Error when attempting to send chat", err)
	}
}
