package models

import (
	"encoding/json"
)

type ChatMessageRequest struct {
	Text string `json:"text"`
}

func (req *ChatMessageRequest) Decode(message *Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}

// ChatMessageResponse Sending to client
type ChatMessageResponse struct {
	Text         string `json:"text"`
	ConnectionId string `json:"connectionId"`
}

func (res ChatMessageResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: Service.Chat, Action: Chat.ServerActions.Receive, Data: string(data)}
	return json.Marshal(message)
}
