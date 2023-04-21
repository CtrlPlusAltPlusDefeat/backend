package models

import (
	"encoding/json"
	"fmt"
)

type ChatMessageRequest struct {
	Text string `json:"text"`
}

func DecodeMessage[Output ChatMessageRequest](message *Wrapper, req *Output) error {
	if fmt.Sprintf("%s/%s", message.Service, message.Action) == "chat/send" {
		return json.Unmarshal([]byte(message.Data), req)
	}
	return nil
}

// ChatMessageResponse Sending to client
type ChatMessageResponse struct {
	Text         string `json:"text"`
	ConnectionId string `json:"connectionId"`
}

func GetChatMessageResponse(connectionId string, text string) ([]byte, error) {
	data, _ := json.Marshal(ChatMessageResponse{Text: text, ConnectionId: connectionId})
	message := Wrapper{Service: "chat", Action: "receive", Data: string(data)}
	return json.Marshal(message)
}
