package services

import (
	"backend/pkg/db"
	game "backend/pkg/game"
	"backend/pkg/game/wordguess"
	"backend/pkg/models"
	"backend/pkg/models/context"
	"backend/pkg/ws"
	"github.com/google/uuid"
	"log"
)

func CreateLobby(context *context.Context, data *models.Data) error {
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

func JoinLobby(context *context.Context, data *models.Data) error {
	req := models.CreateAndJoinRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	return join(context, req.Name, false)
}

func join(context *context.Context, name string, isAdmin bool) error {
	player, err := db.LobbyPlayer.Add(context.LobbyId(), context.SessionId(), context.ConnectionId(), name, isAdmin)

	if err != nil {
		return err
	}

	players, err := db.LobbyPlayer.GetPlayers(context.LobbyId())

	if err != nil {
		return err
	}

	err = ws.Send(context, models.JoinedLobby(), models.GetResponse{Player: player, Lobby: models.Details{Players: players, LobbyId: *context.LobbyId(), Settings: context.Lobby().Settings, InGame: context.Lobby().InGame, GameSessionId: context.Lobby().GameSessionId}})

	if err != nil {
		return err
	}

	return ws.SendToLobby(context, models.PlayerJoined(), models.PlayerJoinResponse{Player: player})
}

func LeaveLobby(context *context.Context, data *models.Data) error {
	player, err := db.LobbyPlayer.UpdateOnline(context.LobbyId(), context.SessionId(), false)

	if err != nil {
		return err
	}

	return ws.SendToLobby(context, models.PlayerLeave(), models.PlayerLeftResponse{Player: player})
}

// LoadGame - This function is called when the host clicks the start game button. It will move the lobby into the selected game in prematch state
func LoadGame(context *context.Context, data *models.Data) error {
	log.Printf("Starting game for lobby '%s'", *context.LobbyId())
	players, err := db.LobbyPlayer.GetPlayers(&context.Lobby().LobbyId)
	if err != nil {
		return err
	}

	teams, err := RandomlyAssignTeams(context.Lobby(), players)
	if err != nil {
		return err
	}

	lobbySettings, err := context.Lobby().Settings.Decode()
	if err != nil {
		return err
	}

	var gState []byte
	switch models.Id(context.GameId()) {
	case models.WordGuess:
		teams, gState, err = wordguess.Setup(teams, lobbySettings)
	}

	gameSession, err := db.GameSession.Add(&game.Session{
		Info: &game.SessionInfo{
			LobbyId:       *context.LobbyId(),
			GameSessionId: uuid.New().String(),
			GameTypeId:    lobbySettings.GameId,
		},
		State: game.NewGameState(),
		Teams: teams,
		Game:  gState,
	})
	if err != nil {
		return err
	}

	updateLobby := models.Lobby{
		LobbyId:       *context.LobbyId(),
		Settings:      context.Lobby().Settings,
		InGame:        true,
		GameSessionId: gameSession.Info.GameSessionId,
	}

	err = db.Lobby.Update(updateLobby)
	if err != nil {
		return err
	}

	return ws.SendToLobby(context, context.Route(), updateLobby)
}
