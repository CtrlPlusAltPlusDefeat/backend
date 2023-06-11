package game

import (
	"backend/pkg/models"
	"encoding/json"
)

func NewGameState(teams TeamArray) *SessionState {
	return &SessionState{
		Teams:       teams,
		CurrentTurn: models.None,
		State:       models.PreMatch,
	}
}

func (e *EncodedGameState) Decode() (*SessionState, error) {
	gameState := SessionState{}
	err := json.Unmarshal([]byte(*e), &gameState)
	if err != nil {
		return nil, err
	}
	return &gameState, nil
}

func (g *SessionState) Encode() (*EncodedGameState, error) {
	marshal, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	encoded := EncodedGameState(marshal)
	return &encoded, nil
}

func (g *SessionState) SetNextTurn() *SessionState {
	currentTurnIndex := g.Teams.GetIndex(g.CurrentTurn)
	if currentTurnIndex == -1 {
		return g
	}

	if currentTurnIndex == len(g.Teams)-1 {
		g.CurrentTurn = g.Teams[0].Name
	} else {
		g.CurrentTurn = g.Teams[currentTurnIndex+1].Name
	}
	return g
}

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
