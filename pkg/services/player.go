package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/ws"
	"context"
	"github.com/google/uuid"
)

type playerT struct {
}

var Player playerT

func (Player playerT) CreateSession(connectionId string) error {
	return Player.SetSession(uuid.New().String(), connectionId)
}

func (Player playerT) SetSession(sessionId string, connectionId string) error {
	connection, err := db.Connection.GetClient().Get(connectionId)
	if err != nil {
		return err
	}
	//check if we need to update the session
	if connection.SessionId != sessionId {
		connection.SessionId = sessionId
		return db.Connection.GetClient().Update(connectionId, db.ConnectionUpdate{SessionId: sessionId})
	}

	//create response
	msg, err := models.SessionResponse{SessionId: sessionId}.Encode()
	if err != nil {
		return err
	}
	return ws.Send(context.TODO(), connectionId, msg)
}
