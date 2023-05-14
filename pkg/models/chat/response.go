package chat

type MessageResponse struct {
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
	PlayerId  string `json:"playerId"`
}
