package routes

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/context"
	"backend/pkg/ws"
)

func ErrorCommunicateMiddleware(next Handler) Handler {
	return func(context *context.Context, data *models.Data) error {
		err := next(context, data)

		if err != nil {
			_ = ws.Send(context, data.Route(), models.ErrorResponse{Error: "Well that didn't work did it. I'd say try again but it probably won't work then either."})

			return err
		}

		return nil
	}
}

func SessionMiddleware(next Handler) Handler {
	return func(context *context.Context, data *models.Data) error {
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

	return func(context *context.Context, data *models.Data) error {
		req := lobbyId{}
		err := data.DecodeTo(&req)

		if err != nil {
			return err
		}

		res, err := db.Lobby.Get(&req.LobbyId)

		if err != nil {
			return err
		}

		return next(context.ForLobby(&res), data)
	}
}

func GameSessionMiddleware(next Handler) Handler {
	type (
		gameSessionId struct {
			GameSessionId string `json:"gameSessionId"`
			LobbyId       string `json:"lobbyId"`
		}
	)

	return func(context *context.Context, data *models.Data) error {
		req := gameSessionId{}
		err := data.DecodeTo(&req)

		if err != nil {
			return err
		}

		res, err := db.GameSession.Get(&req.LobbyId, &req.GameSessionId)
		context.ForLobby(&models.Lobby{LobbyId: req.LobbyId})

		if err != nil {
			return err
		}

		return next(context.ForGameSession(res), data)
	}
}
