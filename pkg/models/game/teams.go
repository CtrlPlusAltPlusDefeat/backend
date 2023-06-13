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
			Players: make([]string, 0),
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

func (t *Team) IncludesPlayer(id string) bool {
	for _, x := range t.Players {
		if x == id {
			return true
		}
	}

	return false
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
	for i, pId := range t.Players {
		if pId == id {
			t.Players = append(t.Players[:i], t.Players[i+1:]...)
		}
	}
}

func (t *Team) AddPlayer(id string) {
	t.Players = append(t.Players, id)
}

func (t TeamArray) SwapTeam(id string, team models.TeamName) TeamArray {
	for i, x := range t {
		if x.IncludesPlayer(id) {
			t[i].RemovePlayer(id)
		}
	}

	t[t.GetIndex(team)].AddPlayer(id)
	return t
}

func (t TeamArray) GetRandom() Team {
	return t[(rand.Intn(len(t)))]
}
