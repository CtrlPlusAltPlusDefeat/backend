package route

import (
	"backend/pkg/models"
	"backend/pkg/models/chat"
	"backend/pkg/services"
	"log"
)

func SendChat(context *models.Context, data *models.Data) error {
	log.Printf("chatHandle: %s", data.Message.Action)

	message := chat.MessageRequest{}
	err := message.Decode(&data.Message)

	if err != nil {
		return err
	}

	return services.BroadcastMessage(context, message)
}
