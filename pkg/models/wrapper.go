package models

import "encoding/json"

type Wrapper struct {
	Service string `json:"service"`
	Action  string `json:"action"`
	Data    string `json:"data"`
}

func (message *Wrapper) Encode() ([]byte, error) {
	return json.Marshal(message)
}

func (message *Wrapper) Decode(data []byte) error {
	return json.Unmarshal(data, message)
}
