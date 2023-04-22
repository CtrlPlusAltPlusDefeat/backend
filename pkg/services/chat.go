package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/ws"
	"context"
	"fmt"
	"log"
)

func BroadcastMessage(connectionId string, chatMessage models.ChatMessageRequest) error {
	connections, err := db.GetConnectionDb().GetAll()
	if err != nil {
		return err
	}
	response, _ := models.GetChatMessageResponse(connectionId, chatMessage.Text)
	fmt.Println("Sending ", chatMessage.Text, " to all connections")
	for _, con := range connections {
		if con.ConnectionId == connectionId {
			continue
		}
		sendChat(con.ConnectionId, response)
	}
	return nil
}

func sendChat(connectionId string, message []byte) {
	err := ws.Send(context.TODO(), connectionId, message)
	if err != nil {
		//we got an error when sending to a client
		log.Printf("Error sending: %s", err)
	}
}
