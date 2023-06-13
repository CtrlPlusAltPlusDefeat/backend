package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/context"
	"backend/pkg/models/game"
	"backend/pkg/ws"
)

func RandomlyAssignTeams(lobby *models.Lobby, players []models.Player) (game.TeamArray, error) {
	settings, err := lobby.Settings.Decode()
	if err != nil {
		return make([]game.Team, 0), err
	}
	teams := game.CreateTeams(settings.Teams)

	for i := 0; i < len(players); i++ {
		teams[i%len(teams)].Players = append(teams[i%len(teams)].Players, players[i].Id)
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

	err = session.IncrementState(player)

	if err != nil {
		return err
	}

	session, err = db.GameSession.Add(context.GameSession())

	if err != nil {
		return err
	}

	return ws.SendToLobby(context, context.Route(), session.State)
}

func GetState(context *context.Context, data *models.Data) error {
	return ws.Send(context, context.Route(), context.GameSession())
}

func SwapTeam(context *context.Context, data *models.Data) error {
	player, err := db.LobbyPlayer.Get(context.LobbyId(), context.SessionId())
	if err != nil {
		return err
	}
	session := context.GameSession()

	req := models.SwapTeamRequest{}
	err = data.DecodeTo(&req)
	if err != nil {
		return err
	}

	session.Teams.SwapTeam(player.Id, req.Team)
	session, err = db.GameSession.Add(context.GameSession())

	return ws.Send(context, context.Route(), session.Teams)
}
