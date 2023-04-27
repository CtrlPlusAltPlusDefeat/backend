package lobby

import (
	"backend/pkg/models"
	"encoding/json"
)

type JoinRequest struct {
	LobbyId string `json:"lobbyId"`
}

func (req *JoinRequest) Decode(message *models.Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}

type GetRequest struct {
	LobbyId string `json:"lobbyId"`
}

func (req *GetRequest) Decode(message *models.Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}

type SetNameRequest struct {
	Text    string `json:"text"`
	LobbyId string `json:"lobbyId"`
}

func (req *SetNameRequest) Decode(message *models.Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}
