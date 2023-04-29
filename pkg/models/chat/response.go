package chat

type MessageResponse struct {
	Text     string `json:"text"`
	PlayerId string `json:"playerId"`
}
