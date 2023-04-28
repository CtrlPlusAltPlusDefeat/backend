package chat

import (
	"backend/pkg/models"
	"encoding/json"
)

type MessageResponse struct {
	Text         string `json:"text"`
	ConnectionId string `json:"connectionId"`
}

func (res MessageResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := models.Wrapper{Service: models.Service.Chat, Action: Actions.Server.Receive, Data: string(data)}
	return json.Marshal(message)
}
