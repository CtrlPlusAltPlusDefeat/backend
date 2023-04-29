package routes

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/ws"
)

func ErrorCommunicateMiddleware(next Handler) Handler {
	return func(context *models.Context, data *models.Data) error {
		err := next(context, data)

		if err != nil {
			res := models.ErrorResponse{Error: "Something went wrong handling this request."}

			err = ws.Send(context, data.Route(), res)

			return err
		}

		return nil
	}
}

func SessionMiddleware(next Handler) Handler {
	return func(context *models.Context, data *models.Data) error {
		res, err := db.Connection.Get(context.ConnectionId())

		if err != nil {
			return err
		}

		return next(context.ForSession(&res.SessionId), data)
	}
}

func LobbyMiddleware(next Handler) Handler {
	type (
		lobbyId struct {
			LobbyId string `json:"lobbyId"`
		}
	)

	return func(context *models.Context, data *models.Data) error {
		req := lobbyId{}
		err := data.DecodeTo(&req)

		if err != nil {
			return err
		}

		return next(context.ForLobby(&req.LobbyId), data)
	}
}
