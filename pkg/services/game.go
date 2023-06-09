package services

import (
	"backend/pkg/models"
	"backend/pkg/models/game"
	"backend/pkg/models/lobby"
	"backend/pkg/models/player"
	"backend/pkg/ws"
)

func RandomlyAssignTeams(lobby *lobby.Lobby, players []player.Player) ([]game.Team, error) {
	settings, err := lobby.Settings.Decode()
	if err != nil {
		return make([]game.Team, 0), err
	}
	teams := make([]game.Team, settings.Teams)

	for i := 0; i < len(players); i++ {
		teams[i%len(teams)].Players = append(teams[i%len(teams)].Players, players[i].Id)
	}

	for i := range teams {
		teams[i].Name = game.GetTeamName(i)
	}

	return teams, nil
}

func PlayerAction(context *models.Context, data *models.Data) error {
	route := models.NewRoute(&models.Service.Lobby, &lobby.Action.Server.LoadGame)
	return ws.SendToLobby(context, route, game.Session{})
}
