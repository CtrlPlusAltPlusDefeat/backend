package lobby

import (
	"backend/pkg/models/chat"
	"backend/pkg/models/player"
)

type Details struct {
	Chats   []chat.Chat     `json:"chats"`
	Players []player.Player `json:"players"`
	LobbyId string          `json:"lobbyId"`
}

type Lobby struct {
	LobbyId string `json:"lobbyId"`
}
