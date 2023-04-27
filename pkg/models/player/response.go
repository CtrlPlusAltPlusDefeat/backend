package player

import (
	"backend/pkg/models"
	"encoding/json"
)

type SessionResponse struct {
	SessionId string `json:"sessionId"`
}

func (res SessionResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := models.Wrapper{Service: models.Service.Player, Action: Action.Server.SetSession, Data: string(data)}
	return json.Marshal(message)
}
