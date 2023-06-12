package game

import (
	"backend/pkg/models"
	"encoding/json"
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

func (t TeamArray) GetIndex(team models.TeamName) int {
	for i, x := range t {
		if x.Name == team {
			return i
		}
	}
	return -1
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

func (t Team) IncludesPlayer(id string) bool {
	for _, x := range t.Players {
		if x == id {
			return true
		}
	}

	return false
}
