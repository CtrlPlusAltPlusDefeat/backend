package game

import (
	"backend/pkg/models"
	"encoding/json"
)

func NewGameState() *SessionState {
	return &SessionState{
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
