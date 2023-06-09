package state

import (
	"backend/pkg/models/game"
	"encoding/json"
)

type EncodedGameState string

type GameState struct {
	Teams       game.TeamArray `dynamodbav:"Teams" json:"teams"`
	CurrentTurn game.TeamName  `dynamodbav:"CurrentTurn" json:"currentTurn"`
	State       game.State     `dynamodbav:"State" json:"state"`
}

func NewGameState(teams game.TeamArray) *GameState {
	return &GameState{
		Teams:       teams,
		CurrentTurn: game.None,
		State:       game.PreMatch,
	}
}

func (e *EncodedGameState) Decode() (*GameState, error) {
	gameState := GameState{}
	err := json.Unmarshal([]byte(*e), &gameState)
	if err != nil {
		return nil, err
	}
	return &gameState, nil
}

func (g *GameState) Encode() (*EncodedGameState, error) {
	marshal, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	encoded := EncodedGameState(marshal)
	return &encoded, nil
}

func (g *GameState) SetNextTurn() *GameState {
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
