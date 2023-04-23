package models

import (
	"encoding/json"
)

// ErrorResponse Sending to client
type ErrorResponse struct {
	Error string `json:"error"`
}

func (res ErrorResponse) New(s string, a string) ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: s, Action: a, Data: string(data)}
	return json.Marshal(message)
}

func (res ErrorResponse) UseWrapper(w Wrapper) ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: w.Service, Action: w.Action, Data: string(data)}
	return json.Marshal(message)
}