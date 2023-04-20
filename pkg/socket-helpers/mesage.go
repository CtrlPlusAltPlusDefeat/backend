package socket_helpers

import "encoding/json"

type Message struct {
	Service string `json:"service"`
	Action  string `json:"action"`
	Data    string `json:"data"`
}

func (message *Message) Encode() ([]byte, error) {
	return json.Marshal(message)
}

func (message *Message) Decode(data []byte) error {
	return json.Unmarshal(data, message)
}
