package services

import (
	"backend/pkg/models"
	"backend/pkg/models/game"
	"backend/pkg/ws"
)

func RandomlyAssignTeams(lobby *models.Lobby, players []models.Player) ([]game.Team, error) {
	settings, err := lobby.Settings.Decode()
	if err != nil {
		return make([]game.Team, 0), err
	}
	teams := make([]game.Team, settings.Teams)

	for i := 0; i < len(players); i++ {
		teams[i%len(teams)].Players = append(teams[i%len(teams)].Players, players[i].Id)
	}

	for i := range teams {
		teams[i].Name = models.GetTeamName(i)
	}

	return teams, nil
}

func PlayerAction(context *models.Context, data *models.Data) error {
	return ws.SendToLobby(context, context.Route(), game.Session{})
}
