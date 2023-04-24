package models

import (
	"encoding/json"
)

type LobbyJoinRequest struct {
	LobbyId string `json:"sessionId"`
}

func (req *LobbyJoinRequest) Decode(message *Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}

type LobbyPlayer struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	IsAdmin bool   `json:"isAdmin"`
	Points  int32  `json:"points"`
}

type LobbyDetails struct {
	Players []LobbyPlayer `json:"players"`
	LobbyId string        `json:"lobbyId"`
}

// LobbyJoinResponse Sending to client
type LobbyJoinResponse struct {
	Player LobbyPlayer  `json:"players"`
	Lobby  LobbyDetails `json:"lobby"`
}

func (res LobbyJoinResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: Service.Lobby, Action: Lobby.ServerActions.Joined, Data: string(data)}
	return json.Marshal(message)
}
