package game

import (
	"backend/pkg/models"
	"encoding/json"
	"math/rand"
)

func CreateTeams(number int) TeamArray {
	var teams TeamArray
	for i := 0; i < number; i++ {
		teams = append(teams, Team{
			Name:    models.GetTeamName(i),
			Players: make([]TeamPlayer, 0),
		})
	}
	return teams
}

func (t TeamArray) Encode() (*EncodedTeamArray, error) {
	marshaled, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	encoded := EncodedTeamArray(marshaled)
	return &encoded, err
}

func (t *EncodedTeamArray) Decode() *TeamArray {
	teamArray := TeamArray{}
	err := json.Unmarshal([]byte(*t), &teamArray)
	if err != nil {
		return nil
	}
	return &teamArray
}

func (t *Team) GetPlayerIndex(id string) int {
	for i, x := range t.Players {
		if x.Id == id {
			return i
		}
	}

	return -1
}

func (t *Team) GetPlayer(id string) *TeamPlayer {
	for _, x := range t.Players {
		if x.Id == id {
			return &x
		}
	}

	return nil
}

func (t TeamArray) GetIndex(team models.TeamName) int {
	for i, x := range t {
		if x.Name == team {
			return i
		}
	}
	return -1
}

func (t *Team) RemovePlayer(id string) {
	for i, p := range t.Players {
		if p.Id == id {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
		}
	}
}

func (t *Team) AddPlayer(p TeamPlayer) {
	t.Players = append(t.Players, p)
}

func (t TeamArray) SwapTeam(id string, team models.TeamName) TeamArray {
	var p *TeamPlayer
	for i, x := range t {
		p = x.GetPlayer(id)
		if p != nil {
			t[i].RemovePlayer(id)
			break
		}
	}

	t[t.GetIndex(team)].AddPlayer(*p)
	return t
}

func (t TeamArray) GetRandom() Team {
	return t[(rand.Intn(len(t)))]
}

func (p *TeamPlayer) DecodeTo(req interface{}) error {
	return json.Unmarshal(p.Data, req)
}
