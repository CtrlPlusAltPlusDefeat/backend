package models

import (
	"encoding/json"
)

type LobbyJoinRequest struct {
	LobbyId string `json:"lobbyId"`
}

func (req *LobbyJoinRequest) Decode(message *Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}

type LobbyGetRequest struct {
	LobbyId string `json:"lobbyId"`
}

func (req *LobbyGetRequest) Decode(message *Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}

type LobbySetNameRequest struct {
	Text    string `json:"text"`
	LobbyId string `json:"lobbyId"`
}

func (req *LobbySetNameRequest) Decode(message *Wrapper) error {
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
	LobbyId string `json:"lobbyId"`
}

func (res LobbyJoinResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: Service.Lobby, Action: Lobby.ServerActions.Joined, Data: string(data)}
	return json.Marshal(message)
}

// LobbyGetResponse Sending to client
type LobbyGetResponse struct {
	Lobby  LobbyDetails `json:"lobby"`
	Player LobbyPlayer  `json:"player"`
}

func (res LobbyGetResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: Service.Lobby, Action: Lobby.ServerActions.Joined, Data: string(data)}
	return json.Marshal(message)
}

// LobbyPlayerJoinResponse Sending to client
type LobbyPlayerJoinResponse struct {
	Player LobbyPlayer `json:"player"`
}

func (res LobbyPlayerJoinResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: Service.Lobby, Action: Lobby.ServerActions.Joined, Data: string(data)}
	return json.Marshal(message)
}

// LobbyPlayerLeftResponse Sending to client
type LobbyPlayerLeftResponse struct {
	Player LobbyPlayer `json:"player"`
}

func (res LobbyPlayerLeftResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: Service.Lobby, Action: Lobby.ServerActions.Joined, Data: string(data)}
	return json.Marshal(message)
}

// LobbyNameChangeResponse Sending to client
type LobbyNameChangeResponse struct {
	Player LobbyPlayer `json:"player"`
}

func (res LobbyNameChangeResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: Service.Lobby, Action: Lobby.ServerActions.Joined, Data: string(data)}
	return json.Marshal(message)
}
