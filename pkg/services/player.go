package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/context"
	"backend/pkg/ws"
	"github.com/google/uuid"
	"log"
)

func CreateSession(context *context.Context, data *models.Data) error {
	return createSession(context)
}

func UseSession(context *context.Context, data *models.Data) error {
	req := models.SessionUseRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	_, err = uuid.Parse(req.SessionId)

	if err != nil {
		return err
	}

	return setSession(context.ForSession(&req.SessionId))
}

func createSession(context *context.Context) error {
	id := uuid.New().String()

	return setSession(context.ForSession(&id))
}

func setSession(context *context.Context) error {
	log.Printf("SetSession for connection %s to %s", *context.ConnectionId(), *context.SessionId())

	connection, err := db.Connection.Get(context.ConnectionId())

	if err != nil {
		return err
	}

	// Delete all sessions using this sessionId
	err = destroySession(context)

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
	res := models.SessionResponse{SessionId: *context.SessionId()}
	return ws.Send(context, models.SetSession(), res)
}

func destroySession(context *context.Context) error {
	connections, err := db.Connection.GetBySessionId(context.SessionId())

	if err != nil || len(connections) == 0 {
		return err
	}

	log.Printf("Destroying %d sessions", len(connections))

	for _, connection := range connections {
		_ = db.Connection.Remove(&connection.ConnectionId)
	}

	return nil
}
