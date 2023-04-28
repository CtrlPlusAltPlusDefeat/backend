package services

import (
	"backend/pkg/db"
	"backend/pkg/models/chat"
	"backend/pkg/ws"
	"context"
	"log"
)

type chatT struct {
}

var Chat chatT

func (c chatT) BroadcastMessage(connectionId string, chatMessage chat.MessageRequest) error {
	connections, err := db.Connection.GetAll()
	if err != nil {
		return err
	}

	response, err := chat.MessageResponse{Text: chatMessage.Text, ConnectionId: connectionId}.Encode()
	if err != nil {
		return err
	}

	log.Println("Sending ", chatMessage.Text, " to ", len(connections), " connections")

	for index, con := range connections {
		log.Println("Sending ", chatMessage.Text, " to connection ", index)

		sendChat(con.ConnectionId, response)
	}
	return nil
}

func sendChat(connectionId string, message []byte) {
	err := ws.Send(context.TODO(), &connectionId, message)
	if err != nil {
		//we got an error when sending to a client
		log.Printf("Error sending: %s", err)
	}
}
