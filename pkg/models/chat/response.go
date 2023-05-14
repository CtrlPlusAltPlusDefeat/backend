package chat

type MessageResponse struct {
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
	PlayerId  string `json:"playerId"`
}
