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

	response, err := chat.MessageResponse{Text: connectionId, ConnectionId: chatMessage.Text}.Encode()
	if err != nil {
		return err
	}

	log.Println("Sending ", chatMessage.Text, " to all connections")
	for _, con := range connections {
		if con.ConnectionId == connectionId {
			continue
		}
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
