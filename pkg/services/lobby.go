package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/game"
	"backend/pkg/models/game/state"
	"backend/pkg/models/lobby"
	"backend/pkg/ws"
	"github.com/google/uuid"
	"log"
)

func CreateLobby(context *models.Context, data *models.Data) error {
	req := lobby.CreateAndJoinRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	id := uuid.New().String()

	context = context.ForLobby(&lobby.Lobby{LobbyId: id})

	err = db.Lobby.Add(context.LobbyId())

	if err != nil {
		return err
	}

	return join(context, req.Name, true)
}

func JoinLobby(context *models.Context, data *models.Data) error {
	req := lobby.CreateAndJoinRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	return join(context, req.Name, false)
}

func join(context *models.Context, name string, isAdmin bool) error {
	player, err := db.LobbyPlayer.Add(context.LobbyId(), context.SessionId(), context.ConnectionId(), name, isAdmin)

	if err != nil {
		return err
	}

	players, err := db.LobbyPlayer.GetPlayers(context.LobbyId())

	if err != nil {
		return err
	}

	route := models.NewRoute(&models.Service.Lobby, &lobby.Action.Server.Joined)
	err = ws.Send(context, route, lobby.GetResponse{Player: player, Lobby: lobby.Details{Players: players, LobbyId: *context.LobbyId(), Settings: context.Lobby().Settings, InGame: context.Lobby().InGame, GameId: context.Lobby().GameId}})

	if err != nil {
		return err
	}

	route = models.NewRoute(&models.Service.Lobby, &lobby.Action.Server.PlayerJoined)
	return ws.SendToLobby(context, route, lobby.PlayerJoinResponse{Player: player})
}

func LeaveLobby(context *models.Context, data *models.Data) error {
	player, err := db.LobbyPlayer.UpdateOnline(context.LobbyId(), context.SessionId(), false)

	if err != nil {
		return err
	}

	route := models.NewRoute(&models.Service.Lobby, &lobby.Action.Server.PlayerLeft)
	return ws.SendToLobby(context, route, lobby.PlayerLeftResponse{Player: player})
}

// LoadGame - This function is called when the host clicks the start game button. It will move the lobby into the selected game in prematch state
func LoadGame(context *models.Context, data *models.Data) error {
	log.Printf("Starting game for lobby '%s'", *context.LobbyId())

	settings, err := context.Lobby().Settings.Decode()
	if err != nil {
		return err
	}

	players, err := db.LobbyPlayer.GetPlayers(&context.Lobby().LobbyId)

	teams, err := RandomlyAssignTeams(context.Lobby(), players)

	gameSession, err := db.GameSession.Add(&game.Session{
		LobbyId:       *context.LobbyId(),
		GameSessionId: uuid.New().String(),
		GameTypeId:    settings.GameId,
		GameState:     state.NewGameState(teams),
	})
	if err != nil {
		return err
	}

	updateLobby := lobby.Lobby{
		LobbyId:  *context.LobbyId(),
		Settings: context.Lobby().Settings,
		InGame:   true,
		GameId:   gameSession.GameSessionId,
	}

	err = db.Lobby.Update(updateLobby)
	if err != nil {
		return err
	}

	route := models.NewRoute(&models.Service.Lobby, &lobby.Action.Server.LoadGame)
	return ws.SendToLobby(context, route, updateLobby)
}
