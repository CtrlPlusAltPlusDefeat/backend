package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/lobby"
	"backend/pkg/ws"
	"github.com/google/uuid"
)

func CreateLobby(context *models.Context, data *models.Data) error {
	id := uuid.New().String()

	context = context.ForLobby(&id)

	err := db.Lobby.Add(context.LobbyId())

	if err != nil {
		return err
	}

	return join(context, true)
}

func JoinLobby(context *models.Context, data *models.Data) error {
	return join(context, false)
}

func SetLobbyName(context *models.Context, data *models.Data) error {
	req := lobby.SetNameRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	player, err := db.LobbyPlayer.UpdateName(context.LobbyId(), context.SessionId(), &req.Text)

	if err != nil {
		return err
	}

	route := models.NewRoute(&models.Service.Lobby, &lobby.Action.Server.NameChanged)
	return ws.SendToLobby(context, route, lobby.NameChangeResponse{Player: player})
}

func join(context *models.Context, isAdmin bool) error {
	player, err := db.LobbyPlayer.Get(context.LobbyId(), context.SessionId())

	if err != nil {
		player, err = db.LobbyPlayer.Add(context.LobbyId(), context.SessionId(), context.ConnectionId(), isAdmin)
	}

	if err != nil {
		return err
	}

	route := models.NewRoute(&models.Service.Lobby, &lobby.Action.Server.PlayerJoined)
	err = ws.SendToLobby(context, route, lobby.PlayerJoinResponse{Player: player})

	if err != nil {
		return err
	}

	players, err := db.LobbyPlayer.GetPlayers(context.LobbyId())

	if err != nil {
		return err
	}

	route = models.NewRoute(&models.Service.Lobby, &lobby.Action.Server.Joined)
	return ws.Send(context, route, lobby.GetResponse{Player: player, Lobby: lobby.Details{Players: players, LobbyId: *context.LobbyId()}})
}
