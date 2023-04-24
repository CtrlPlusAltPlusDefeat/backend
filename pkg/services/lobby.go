package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/ws"
	"context"
	"github.com/google/uuid"
	"log"
)

type lobbyT struct {
}

var Lobby lobbyT

func (s lobbyT) Create() error {
	lobbyId := uuid.New().String()
	err := db.Lobby.Add(lobbyId)
	if err != nil {
		return err
	}
	return s.Join(lobbyId, true)
}

func (s lobbyT) Join(lobbyId string, isAdmin bool) error {
	player, err := db.LobbyPlayer.Add(&lobbyId, SocketData.SessionId, &SocketData.RequestContext.ConnectionID, isAdmin)
	if err != nil {
		return err
	}
	err = onPlayerJoin(&player, &lobbyId)
	if err != nil {
		return err
	}
	return sendLobbyJoin(&lobbyId)
}

func (s lobbyT) Get(lobbyId *string) error {
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
		if p.SessionId == *SocketData.SessionId {
			thisPlayer = player
		}
		allPlayers = append(allPlayers, player)
	}

	res := models.LobbyGetResponse{Player: thisPlayer, Lobby: models.LobbyDetails{
		Players: allPlayers,
		LobbyId: *lobbyId,
	}}
	bytes, err := res.Encode()
	if err != nil {
		return err
	}

	return ws.Send(context.TODO(), &SocketData.RequestContext.ConnectionID, bytes)
}

func (s lobbyT) NameChange(name *string, lobbyId *string) error {
	player, err := db.LobbyPlayer.UpdateName(lobbyId, SocketData.SessionId, name)
	if err != nil {
		return err
	}
	return onPlayerNameChange(&player, lobbyId)
}

func sendLobbyJoin(lobbyId *string) error {
	res := models.LobbyJoinResponse{LobbyId: *lobbyId}
	bytes, err := res.Encode()
	if err != nil {
		return err
	}
	return ws.Send(context.TODO(), &SocketData.RequestContext.ConnectionID, bytes)
}

func onPlayerJoin(player *db.Player, lobbyId *string) error {

	response := models.LobbyPlayerJoinResponse{Player: models.LobbyPlayer{
		Id:      player.Id,
		Name:    player.Name,
		IsAdmin: player.IsAdmin,
		Points:  player.Points,
	}}
	bytes, err := response.Encode()
	if err != nil {
		log.Printf("onPlayerJoin error encoding response")
		return err
	}
	return sendToLobby(lobbyId, bytes)
}

func onPlayerNameChange(player *db.Player, lobbyId *string) error {
	response := models.LobbyNameChangeResponse{Player: models.LobbyPlayer{
		Id:      player.Id,
		Name:    player.Name,
		IsAdmin: player.IsAdmin,
		Points:  player.Points,
	}}
	bytes, err := response.Encode()
	if err != nil {
		log.Printf("onPlayerNameChange error encoding response")
		return err
	}
	return sendToLobby(lobbyId, bytes)
}

func sendToLobby(lobbyId *string, msg []byte) error {
	players, err := db.LobbyPlayer.GetPlayers(lobbyId)
	if err != nil {
		return err
	}
	for _, p := range players {
		err = ws.Send(context.TODO(), &p.ConnectionId, msg)
		if err != nil {
			log.Printf("sendToLobby error sending to %s ", p.ConnectionId)
		}
	}
	return nil
}
