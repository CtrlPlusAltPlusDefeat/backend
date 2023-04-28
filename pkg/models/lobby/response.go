package lobby

import (
	"backend/pkg/models"
	"encoding/json"
)

type JoinResponse struct {
	LobbyId string `json:"lobbyId"`
}

func (res JoinResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := models.Wrapper{Service: models.Service.Lobby, Action: Action.Server.Joined, Data: string(data)}
	return json.Marshal(message)
}

type GetResponse struct {
	Lobby  Details `json:"lobby"`
	Player Player  `json:"player"`
}

func (res GetResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := models.Wrapper{Service: models.Service.Lobby, Action: Action.Server.Get, Data: string(data)}
	return json.Marshal(message)
}

type PlayerJoinResponse struct {
	Player Player `json:"player"`
}

func (res PlayerJoinResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := models.Wrapper{Service: models.Service.Lobby, Action: Action.Server.PlayerJoined, Data: string(data)}
	return json.Marshal(message)
}

type PlayerLeftResponse struct {
	Player Player `json:"player"`
}

func (res PlayerLeftResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := models.Wrapper{Service: models.Service.Lobby, Action: Action.Server.PlayerLeft, Data: string(data)}
	return json.Marshal(message)
}

type NameChangeResponse struct {
	Player Player `json:"player"`
}

func (res NameChangeResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := models.Wrapper{Service: models.Service.Lobby, Action: Action.Server.NameChanged, Data: string(data)}
	return json.Marshal(message)
}
