package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/chat"
	"backend/pkg/ws"
	"log"
)

func BroadcastMessage(context *models.Context, chatMessage chat.MessageRequest) error {
	connections, err := db.Connection.GetAll()
	if err != nil {
		return err
	}

	response, err := chat.MessageResponse{Text: chatMessage.Text, ConnectionId: *context.ConnectionId()}.Encode()
	if err != nil {
		return err
	}

	log.Println("Sending ", chatMessage.Text, " to ", len(connections), " connections")

	for index, con := range connections {
		log.Println("Sending ", chatMessage.Text, " to connection ", index)

		sendChat(context.ForConnection(&con.ConnectionId), response)
	}
	return nil
}

func sendChat(context *models.Context, message []byte) {
	err := ws.Send(context, message)

	if err != nil {
		log.Printf("Error sending: %s", err)
	}
}
