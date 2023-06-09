package lobby

import (
	"backend/pkg/models/game/settings"
	"backend/pkg/models/player"
)

type Details struct {
	Players  []player.Player  `json:"players"`
	LobbyId  string           `json:"lobbyId"`
	Settings settings.Encoded `json:"settings"`
	InGame   bool             `json:"inGame"`
	GameId   string           `json:"gameId"`
}

type Lobby struct {
	LobbyId  string           `json:"lobbyId" dynamodbav:"LobbyId"`
	Settings settings.Encoded `json:"settings" dynamodbav:"Settings"`
	InGame   bool             `json:"inGame" dynamodbav:"InGame"`
	GameId   string           `json:"gameId" dynamodbav:"GameId"`
}
