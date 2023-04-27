package services

import (
	"backend/pkg/db"
	"backend/pkg/models/lobby"
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

	// else add new lobby player
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
	// check if connectionId has changed
	player, err := db.LobbyPlayer.Get(lobbyId, SocketData.SessionId)
	if err != nil {
		return err
	}

	// if so, update connectionId
	if player.ConnectionId == SocketData.RequestContext.ConnectionID {
		player, err = db.LobbyPlayer.UpdateConnectionId(lobbyId, SocketData.SessionId, &SocketData.RequestContext.ConnectionID)
		if err != nil {
			return err
		}
	}

	players, err := db.LobbyPlayer.GetPlayers(lobbyId)
	if err != nil {
		return err
	}

	var thisPlayer lobby.Player
	for _, p := range players {
		if p.SessionId == *SocketData.SessionId {
			thisPlayer = p
			break
		}
	}

	res := lobby.GetResponse{Player: thisPlayer, Lobby: lobby.Details{
		Players: players,
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
	res := lobby.JoinResponse{LobbyId: *lobbyId}
	bytes, err := res.Encode()
	if err != nil {
		return err
	}
	return ws.Send(context.TODO(), &SocketData.RequestContext.ConnectionID, bytes)
}

func onPlayerJoin(player *lobby.Player, lobbyId *string) error {

	response := lobby.PlayerJoinResponse{Player: *player}
	bytes, err := response.Encode()
	if err != nil {
		log.Printf("onPlayerJoin error encoding response")
		return err
	}
	//we don't need to notify this connection as we know we are in the lobby
	return ws.SendToLobby(lobbyId, bytes, true)
}

func onPlayerNameChange(player *lobby.Player, lobbyId *string) error {
	response := lobby.NameChangeResponse{Player: *player}
	bytes, err := response.Encode()
	if err != nil {
		log.Printf("onPlayerNameChange error encoding response")
		return err
	}
	return ws.SendToLobby(lobbyId, bytes, false)
}
