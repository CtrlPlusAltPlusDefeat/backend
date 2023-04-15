package ws

import "encoding/json"

type Message struct {
	Service string          `json:"service"`
	Action  string          `json:"action"`
	Data    json.RawMessage `json:"data"`
}

type ChatMessage struct {
	Text         string `json:"text"`
	ConnectionId string `json:"connectionId"`
}

func (message *Message) Decode(data []byte) (*Message, error) {
	e := json.Unmarshal(data, message)
	return message, e
}
func (message *Message) Encode() ([]byte, error) {
	return json.Marshal(message)
}

func (message *Message) EncodeChatMessage(connectionId string, text string) ([]byte, error) {
	message.Service = "chat"
	message.Action = "received"
	chatMessage, _ := json.Marshal(ChatMessage{Text: text, ConnectionId: connectionId})
	message.Data = chatMessage
	return json.Marshal(message)
}
