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
	settings, err := lobby.Settings.Decode()
	if err != nil {
		return make([]game.Team, 0), err
	}
	teams := game.CreateTeams(settings.Teams)

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
	session := context.GameSession()

	ctx := game.NewContext(context.Value(), session, data, &player)
	switch models.Id(context.GameId()) {
	case models.WordGuess:
		session, err = wordguess.HandlePlayerAction(ctx)
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

func GetState(context *context.Context, data *models.Data) error {
	return ws.Send(context, context.Route(), context.GameSession())
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

	ctx := game.NewContext(context.Value(), session, data, &player)
	switch models.Id(context.GameId()) {
	case models.WordGuess:
		session.Teams, err = wordguess.HandleSwapTeam(ctx)
		break
	}
	if err != nil {
		return err
	}

	session, err = db.GameSession.Add(session)
	return ws.SendToLobby(context, context.Route(), session.Teams)
}
