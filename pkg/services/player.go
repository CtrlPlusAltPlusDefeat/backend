package services

import (
	"backend/pkg/db"
	"backend/pkg/models/player"
	"backend/pkg/ws"
	"context"
	"github.com/google/uuid"
	"log"
)

type playerT struct {
}

var Player playerT

func (Player playerT) CreateSession(connectionId string) error {
	return Player.SetSession(uuid.New().String(), connectionId)
}

func (Player playerT) SetSession(sessionId string, connectionId string) error {
	log.Printf("SetSession for connection %s to %s", connectionId, sessionId)
	connection, err := db.Connection.Get(connectionId)
	if err != nil {
		return err
	}

	// Delete all sessions using this sessionId
	err = DestroySession(sessionId)
	if err != nil {
		return err
	}

	//check if we need to update the session
	if connection.SessionId != sessionId {

		connection.SessionId = sessionId
		err = db.Connection.Update(connectionId, sessionId)
		if err != nil {
			return err
		}
	}

	//create response
	msg, err := player.SessionResponse{SessionId: sessionId}.Encode()
	if err != nil {
		return err
	}
	return ws.Send(context.TODO(), &connectionId, msg)
}

func DestroySession(sessionId string) error {
	connections, err := db.Connection.GetBySessionId(sessionId)
	if err != nil || len(connections) == 0 {
		return err
	}
	log.Printf("Destroying %d sessions", len(connections))
	for _, connection := range connections {
		_ = ws.Disconnect(&connection.ConnectionId)
	}
	return nil
}
