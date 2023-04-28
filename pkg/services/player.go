package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/player"
	"backend/pkg/ws"
	"github.com/google/uuid"
	"log"
)

func CreateSession(context *models.Context) error {
	id := uuid.New().String()

	return SetSession(context.ForSession(&id))
}

func SetSession(context *models.Context) error {
	log.Printf("SetSession for connection %s to %s", *context.ConnectionId(), *context.SessionId())

	connection, err := db.Connection.Get(context.ConnectionId())

	if err != nil {
		return err
	}

	// Delete all sessions using this sessionId
	err = DestroySession(context)

	if err != nil {
		return err
	}

	//check if we need to update the session
	if connection.SessionId != *context.SessionId() {

		err = db.Connection.Update(context.ConnectionId(), context.SessionId())

		if err != nil {
			return err
		}
	}

	//create response
	msg, err := player.SessionResponse{SessionId: *context.SessionId()}.Encode()

	if err != nil {
		return err
	}

	return ws.Send(context, msg)
}

func DestroySession(context *models.Context) error {
	connections, err := db.Connection.GetBySessionId(context.SessionId())

	if err != nil || len(connections) == 0 {
		return err
	}

	log.Printf("Destroying %d sessions", len(connections))

	for _, connection := range connections {
		_ = ws.Disconnect(&connection.ConnectionId)
	}

	return nil
}
