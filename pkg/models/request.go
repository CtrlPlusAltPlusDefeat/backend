package models

type SendChatRequest struct {
	Text string `json:"text"`
}

type LoadChatRequest struct {
	Timestamp int64 `json:"timestamp"`
}

type SessionUseRequest struct {
	SessionId string `json:"sessionId"`
}

type CreateAndJoinRequest struct {
	Name string `json:"name"`
}
