package game

import (
	"backend/pkg/models/game/state"
	"encoding/json"
)

type Session struct {
	LobbyId       string           `dynamodbav:"LobbyId" json:"lobbyId"`
	GameSessionId string           `dynamodbav:"GameSessionId" json:"gameSessionId"`
	GameTypeId    Id               `dynamodbav:"GameTypeId" json:"gameTypeId"`
	GameState     *state.GameState `dynamodbav:"-" json:"gameState"`
}

type Team struct {
	Name    TeamName
	Players []string
}
type TeamArray []Team
type EncodedTeamArray string

func (t TeamArray) GetIndex(team TeamName) int {
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
