package wordguess

import (
	"backend/pkg/game"
	"backend/pkg/models"
	"encoding/json"
)

type role string

const (
	Operative role = "operative"
	SpyMaster role = "spymaster"
)

type PlayerData struct {
	Role role `json:"role"`
}

func AddRoleDefaults(teams game.TeamArray) (game.TeamArray, error) {
	defaultData := PlayerData{
		Role: Operative,
	}

	encoded, err := defaultData.Encode()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(teams); i++ {
		for j := 0; j < len(teams[i].Players); j++ {
			teams[i].Players[j].Data = encoded
		}
	}
	return teams, nil
}

func (p *PlayerData) Encode() ([]byte, error) {
	return json.Marshal(p)
}

func HandleSwapTeam(session *game.Session, data *models.Data, player models.Player) (game.TeamArray, error) {
	var playerData PlayerData
	req := SwapTeamRequest{}
	err := data.DecodeTo(&req)
	if err != nil {
		return nil, err
	}
	teams := session.Teams.SwapTeam(player.Id, req.Team)

	tIndex := teams.GetIndex(req.Team)
	pIndex := teams[tIndex].GetPlayerIndex(player.Id)

	err = teams[tIndex].Players[pIndex].DecodeTo(&playerData)
	if err != nil {
		return nil, err
	}

	playerData.Role = req.Role
	encodedData, err := playerData.Encode()
	if err != nil {
		return nil, err
	}

	teams[tIndex].Players[pIndex].Data = encodedData
	return teams, nil
}
