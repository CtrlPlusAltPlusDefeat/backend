package ws

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"log"
)

var OnLeaveLobby func(context *models.Context, data *models.Data) error

func Disconnect(context *models.Context) error {
	// todo we want to notify lobby that this player has disconnected
	err := OnLeaveLobby(context, nil)
	if err != nil {
		log.Printf("Error notifying lobby: %s", err.Error())
	}
	return deleteConnection(context.ConnectionId())
}

// handle disconnecting connections
func deleteConnection(id *string) error {
	log.Printf("Deleting connection Id: %s", *id)
	return db.Connection.Remove(id)
}
