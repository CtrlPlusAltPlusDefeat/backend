package chat

type SendChatResponse struct {
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
	PlayerId  string `json:"playerId"`
}

type LoadChatResponse struct {
	Messages []SendChatResponse `json:"messages"`
}
