package ws

import "encoding/json"

type Message struct {
	Service string          `json:"service"`
	Action  string          `json:"action"`
	Data    json.RawMessage `json:"data"`
}

func (message *Message) Decode(data []byte) (*Message, error) {
	e := json.Unmarshal(data, message)
	return message, e
}
func (message *Message) Encode() ([]byte, error) {
	return json.Marshal(message)
}


