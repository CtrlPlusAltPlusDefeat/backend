package ws

import (
	"encoding/json"
	"fmt"
)

// Message Generic wrapper for all websocket messages
type Message struct {
	Service string `json:"service"`
	Action  string `json:"action"`
	Data    string `json:"data"`
}

// ChatMessageRequest Received from client
type ChatMessageRequest struct {
	Text string `json:"text"`
}

func (message *Message) Encode() ([]byte, error) {
	return json.Marshal(message)
}

func (message *Message) Decode(data []byte) error {
	return json.Unmarshal(data, message)
}

func DecodeMessage[Output ChatMessageRequest](message *Message, req *Output) error {
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
	message := Message{Service: "chat", Action: "receive", Data: string(data)}
	return json.Marshal(message)
}
