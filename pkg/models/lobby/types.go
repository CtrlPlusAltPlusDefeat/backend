package lobby

import (
	"backend/pkg/models/chat"
	"backend/pkg/models/player"
)

type Details struct {
	Chats    []chat.Chat     `json:"chats"`
	Players  []player.Player `json:"players"`
	LobbyId  string          `json:"lobbyId"`
	Settings string          `json:"settings"`
}

type Lobby struct {
	LobbyId  string `json:"lobbyId" dynamodbav:"LobbyId"`
	Settings string `json:"settings" dynamodbav:"Settings"`
}
