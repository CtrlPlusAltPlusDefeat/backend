package ws

import (
	"backend/pkg/db"
	"log"
)

func Disconnect(id *string) error {
	// todo we want to notify lobby that this player has disconnected

	// todo then remove lobby from their connection
	return deleteConnection(id)
}

// handle disconnecting connections
func deleteConnection(id *string) error {
	log.Printf("Deleting connection Id: %s", id)
	return db.Connection.GetClient().Remove(id)
}
