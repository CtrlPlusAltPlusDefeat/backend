package game

import (
	"backend/pkg/models"
	"encoding/json"
)

type Session struct {
	Info  *SessionInfo    `dynamodbav:"-" json:"info"`
	State *SessionState   `dynamodbav:"-" json:"state"`
	Teams TeamArray       `dynamodbav:"-" json:"teams"`
	Game  json.RawMessage `dynamodbav:"-" json:"game"`
}

type SessionInfo struct {
	LobbyId       string    `dynamodbav:"LobbyId" json:"lobbyId"`
	GameSessionId string    `dynamodbav:"GameSessionId" json:"gameSessionId"`
	GameTypeId    models.Id `dynamodbav:"GameTypeId" json:"gameTypeId"`
}

type SessionState struct {
	CurrentTurn models.TeamName `dynamodbav:"CurrentTurn" json:"currentTurn"`
	State       models.State    `dynamodbav:"State" json:"state"`
}

type EncodedGameState string

type Team struct {
	Name    models.TeamName `json:"name"`
	Players []TeamPlayer    `json:"players"`
}

type TeamPlayer struct {
	Id   string          `json:"id"`
	Data json.RawMessage `json:"data"`
}

type TeamArray []Team

type EncodedTeamArray string
