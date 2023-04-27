package chat

import (
	"backend/pkg/models"
	"encoding/json"
)

type MessageRequest struct {
	Text string `json:"text"`
}

func (req *MessageRequest) Decode(message *models.Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}
