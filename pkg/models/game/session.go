package game

import "backend/pkg/models"

type Session struct {
	LobbyId       string    `dynamodbav:"LobbyId" json:"lobbyId"`
	GameSessionId string    `dynamodbav:"GameSessionId" json:"gameSessionId"`
	GameTypeId    models.Id `dynamodbav:"GameTypeId" json:"gameTypeId"`
	GameState     *State    `dynamodbav:"-" json:"gameState"`
}
