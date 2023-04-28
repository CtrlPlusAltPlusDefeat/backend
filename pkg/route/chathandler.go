package route

import (
	"backend/pkg/models"
	"backend/pkg/models/chat"
	"backend/pkg/services"
	"log"
)

func chatHandle(context *models.Context, data *models.Data) {
	log.Printf("chatHandle: %s", data.Message.Action)

	message := chat.MessageRequest{}
	err := message.Decode(&data.Message)

	if err != nil {
		log.Println("Error decoding message", err)
		return
	}

	err = services.BroadcastMessage(context, message)

	if err != nil {
		log.Println("Error when attempting to send chat", err)
	}
}
