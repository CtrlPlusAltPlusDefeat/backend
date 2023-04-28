package player

import (
	"backend/pkg/models"
	"encoding/json"
)

type SessionUseRequest struct {
	SessionId string `json:"sessionId"`
}

func (req *SessionUseRequest) Decode(message *models.Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}
