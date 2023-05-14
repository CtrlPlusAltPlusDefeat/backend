package chat

type SendChatRequest struct {
	Text string `json:"text"`
}

type LoadChatRequest struct {
	Timestamp int64 `json:"timestamp"`
}
