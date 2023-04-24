package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/ws"
	"context"
	"github.com/google/uuid"
)

type lobbyT struct {
}

var Lobby lobbyT

func (s lobbyT) Create(connectionId *string, sessionId *string) error {
	lobbyId := uuid.New().String()
	err := db.Lobby.Add(lobbyId)
	if err != nil {
		return err
	}
	return s.Join(lobbyId, connectionId, sessionId, true)
}

func (s lobbyT) Join(lobbyId string, connectionId *string, sessionId *string, isAdmin bool) error {
	err := db.LobbyPlayer.Add(&lobbyId, sessionId, isAdmin)
	if err != nil {
		return err
	}
	onPlayerJoin()
	return sendLobbyJoin(connectionId, sessionId, &lobbyId)
}

func sendLobbyJoin(connectionId *string, sessionId *string, lobbyId *string) error {
	playersRes, err := db.LobbyPlayer.GetPlayers(lobbyId)
	if err != nil {
		return err
	}

	var thisPlayer models.LobbyPlayer
	var allPlayers []models.LobbyPlayer
	for _, p := range playersRes {
		player := models.LobbyPlayer{
			Id:      p.Id,
			Name:    p.Name,
			IsAdmin: p.IsAdmin,
			Points:  p.Points,
		}
		if p.SessionId == *sessionId {
			thisPlayer = player
		}
		allPlayers = append(allPlayers, player)
	}

	res := models.LobbyJoinResponse{Player: thisPlayer, Lobby: models.LobbyDetails{
		Players: allPlayers,
		LobbyId: *lobbyId,
	}}
	bytes, err := res.Encode()
	if err != nil {
		return err
	}

	return ws.Send(context.TODO(), connectionId, bytes)
}

func onPlayerJoin() {

}
