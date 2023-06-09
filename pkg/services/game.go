package services

import (
	"backend/pkg/models/game"
	"backend/pkg/models/lobby"
	"backend/pkg/models/player"
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

	return teams, nil
}
