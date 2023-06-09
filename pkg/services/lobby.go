package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/game"
	"backend/pkg/ws"
	"github.com/google/uuid"
	"log"
)

func CreateLobby(context *models.Context, data *models.Data) error {
	req := models.CreateAndJoinRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	id := uuid.New().String()

	context = context.ForLobby(&models.Lobby{LobbyId: id})

	err = db.Lobby.Add(context.LobbyId())

	if err != nil {
		return err
	}

	return join(context, req.Name, true)
}

func JoinLobby(context *models.Context, data *models.Data) error {
	req := models.CreateAndJoinRequest{}
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

	err = ws.Send(context, context.Route(), models.GetResponse{Player: player, Lobby: models.Details{Players: players, LobbyId: *context.LobbyId(), Settings: context.Lobby().Settings, InGame: context.Lobby().InGame, GameId: context.Lobby().GameId}})

	if err != nil {
		return err
	}

	return ws.SendToLobby(context, models.PlayerJoin(), models.PlayerJoinResponse{Player: player})
}

func LeaveLobby(context *models.Context, data *models.Data) error {
	player, err := db.LobbyPlayer.UpdateOnline(context.LobbyId(), context.SessionId(), false)

	if err != nil {
		return err
	}

	return ws.SendToLobby(context, models.PlayerLeave(), models.PlayerLeftResponse{Player: player})
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
		GameState:     game.NewGameState(teams),
	})
	if err != nil {
		return err
	}

	updateLobby := models.Lobby{
		LobbyId:  *context.LobbyId(),
		Settings: context.Lobby().Settings,
		InGame:   true,
		GameId:   gameSession.GameSessionId,
	}

	err = db.Lobby.Update(updateLobby)
	if err != nil {
		return err
	}

	return ws.SendToLobby(context, context.Route(), updateLobby)
}
