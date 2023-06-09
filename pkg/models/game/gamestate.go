package game

import (
	"backend/pkg/models"
	"encoding/json"
)

type EncodedGameState string

type State struct {
	Teams       TeamArray       `dynamodbav:"Teams" json:"teams"`
	CurrentTurn models.TeamName `dynamodbav:"CurrentTurn" json:"currentTurn"`
	State       models.State    `dynamodbav:"State" json:"state"`
}

func NewGameState(teams TeamArray) *State {
	return &State{
		Teams:       teams,
		CurrentTurn: models.None,
		State:       models.PreMatch,
	}
}

func (e *EncodedGameState) Decode() (*State, error) {
	gameState := State{}
	err := json.Unmarshal([]byte(*e), &gameState)
	if err != nil {
		return nil, err
	}
	return &gameState, nil
}

func (g *State) Encode() (*EncodedGameState, error) {
	marshal, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	encoded := EncodedGameState(marshal)
	return &encoded, nil
}

func (g *State) SetNextTurn() *State {
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

type Team struct {
	Name    models.TeamName
	Players []string
}
type TeamArray []Team
type EncodedTeamArray string

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
