package services

import (
	"backend/pkg/db"
	"backend/pkg/game"
	"backend/pkg/game/wordguess"
	"backend/pkg/models"
	"backend/pkg/models/context"
	"backend/pkg/ws"
	"fmt"
)

func RandomlyAssignTeams(lobby *models.Lobby, players []models.Player) (game.TeamArray, error) {
	teams := game.CreateTeams(lobby.Settings.Teams)

	for i := 0; i < len(players); i++ {
		teams[i%len(teams)].Players = append(teams[i%len(teams)].Players, game.TeamPlayer{Id: players[i].Id})
	}

	for i := range teams {
		teams[i].Name = models.GetTeamName(i)
	}

	return teams, nil
}

func PlayerAction(context *context.Context, data *models.Data) error {
	player, err := db.LobbyPlayer.Get(context.LobbyId(), context.SessionId())
	if err != nil {
		return err
	}

	var session *game.Session
	switch models.Id(context.GameId()) {
	case models.WordGuess:
		context, session, err = wordguess.HandlePlayerAction(context, data, &player)
		break
	}
	if err != nil {
		return err
	}

	session, err = db.GameSession.Add(session)
	if err != nil {
		return err
	}

	return ws.SendToLobby(context, context.Route(), session)
}

// GetState this is an expensive call, only use when necessary
func GetState(context *context.Context, data *models.Data) error {
	var session *game.Session
	var err error
	switch models.Id(context.GameId()) {
	case models.WordGuess:
		context, session, err = wordguess.HandleGetState(context, data)
		break
	}
	if err != nil {
		return err
	}
	return ws.Send(context, context.Route(), session)
}

func SwapTeam(context *context.Context, data *models.Data) error {
	session := context.GameSession()
	if session.State.State != models.PreMatch {
		return fmt.Errorf("cannot swap teams when not in %s", models.PreMatch)
	}
	player, err := db.LobbyPlayer.Get(context.LobbyId(), context.SessionId())
	if err != nil {
		return err
	}

	switch models.Id(context.GameId()) {
	case models.WordGuess:
		context, session.Teams, err = wordguess.HandleSwapTeam(context, data, &player)
		break
	}
	if err != nil {
		return err
	}

	session, err = db.GameSession.Add(session)
	return ws.SendToLobby(context, context.Route(), session.Teams)
}
