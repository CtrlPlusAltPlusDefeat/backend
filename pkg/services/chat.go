package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/chat"
	"backend/pkg/ws"
	"log"
)

func SendChat(context *models.Context, data *models.Data) error {
	req := chat.MessageRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	return broadcastMessage(context, req)
}

func broadcastMessage(context *models.Context, chatMessage chat.MessageRequest) error {
	connections, err := db.Connection.GetAll()
	if err != nil {
		return err
	}

	response := chat.MessageResponse{Text: chatMessage.Text, ConnectionId: *context.ConnectionId()}
	route := models.NewRoute(&models.Service.Chat, &chat.Actions.Server.Receive)

	log.Println("Sending ", chatMessage.Text, " to ", len(connections), " connections")

	for index, con := range connections {
		log.Println("Sending ", chatMessage.Text, " to connection ", index)

		err := ws.Send(context.ForConnection(&con.ConnectionId), route, response)

		if err != nil {
			log.Printf("Error sending: %s", err)
		}
	}
	return nil
}
