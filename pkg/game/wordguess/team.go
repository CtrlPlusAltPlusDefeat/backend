package wordguess

import (
	"backend/pkg/game"
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

func HandleSwapTeam(ctx *game.Context) (game.TeamArray, error) {
	var playerData PlayerData
	req := SwapTeamRequest{}
	err := ctx.Data().DecodeTo(&req)
	if err != nil {
		return nil, err
	}
	teams := ctx.Session().Teams.SwapTeam(ctx.Player().Id, req.Team)

	tIndex := teams.GetIndex(req.Team)
	pIndex := teams[tIndex].GetPlayerIndex(ctx.Player().Id)

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
