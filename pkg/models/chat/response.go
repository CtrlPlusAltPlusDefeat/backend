package chat

type MessageResponse struct {
	Text         string `json:"text"`
	ConnectionId string `json:"connectionId"`
}
