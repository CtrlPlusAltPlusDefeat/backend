package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/lobby"
	"backend/pkg/ws"
	"github.com/google/uuid"
	"log"
)

func Create(context *models.Context) error {
	id := uuid.New().String()

	context = context.ForSession(&id)

	err := db.Lobby.Add(context.LobbyId())

	if err != nil {
		return err
	}

	return Join(context, true)
}

func Join(context *models.Context, isAdmin bool) error {
	player, err := db.LobbyPlayer.Add(context.LobbyId(), context.SessionId(), context.ConnectionId(), isAdmin)

	if err != nil {
		return err
	}

	err = onPlayerJoin(context, &player)

	if err != nil {
		return err
	}

	return sendLobbyJoin(context)
}

func Get(context *models.Context) error {
	player, err := db.LobbyPlayer.Get(context.LobbyId(), context.SessionId())

	if err != nil {
		return err
	}

	if player.ConnectionId == *context.ConnectionId() {
		player, err = db.LobbyPlayer.UpdateConnectionId(context.LobbyId(), context.SessionId(), context.ConnectionId())

		if err != nil {
			return err
		}
	}

	players, err := db.LobbyPlayer.GetPlayers(context.LobbyId())

	if err != nil {
		return err
	}

	var thisPlayer lobby.Player
	for _, p := range players {
		if p.SessionId == *context.SessionId() {
			thisPlayer = p
			break
		}
	}

	res := lobby.GetResponse{Player: thisPlayer, Lobby: lobby.Details{
		Players: players,
		LobbyId: *context.LobbyId(),
	}}

	bytes, err := res.Encode()

	if err != nil {
		return err
	}

	return ws.Send(context, bytes)
}

func NameChange(context *models.Context, name *string) error {
	player, err := db.LobbyPlayer.UpdateName(context.LobbyId(), context.SessionId(), name)

	if err != nil {
		return err
	}

	return onPlayerNameChange(context, &player)
}

func sendLobbyJoin(context *models.Context) error {
	res := lobby.JoinResponse{LobbyId: *context.LobbyId()}
	bytes, err := res.Encode()

	if err != nil {
		return err
	}

	return ws.Send(context, bytes)
}

func onPlayerJoin(context *models.Context, player *lobby.Player) error {
	response := lobby.PlayerJoinResponse{Player: *player}
	bytes, err := response.Encode()

	if err != nil {
		log.Printf("onPlayerJoin error encoding response")
		return err
	}

	return ws.SendToLobby(context, bytes, true)
}

func onPlayerNameChange(context *models.Context, player *lobby.Player) error {
	response := lobby.NameChangeResponse{Player: *player}
	bytes, err := response.Encode()

	if err != nil {
		log.Printf("onPlayerNameChange error encoding response")
		return err
	}

	return ws.SendToLobby(context, bytes, false)
}
