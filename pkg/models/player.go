package models

import (
	"encoding/json"
)

type SessionUseRequest struct {
	SessionId string `json:"sessionId"`
}

func (req *SessionUseRequest) Decode(message *Wrapper) error {
	return json.Unmarshal([]byte(message.Data), req)
}

// SessionResponse Sending to client
type SessionResponse struct {
	SessionId string `json:"sessionId"`
}

func (res SessionResponse) Encode() ([]byte, error) {
	data, _ := json.Marshal(res)
	message := Wrapper{Service: Service.Player, Action: Player.ServerActions.SetSession, Data: string(data)}
	return json.Marshal(message)
}
