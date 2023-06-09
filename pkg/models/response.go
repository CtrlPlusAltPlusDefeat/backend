package models

type Wrapper struct {
	Service string `json:"service"`
	Action  string `json:"action"`
	Data    string `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SessionResponse struct {
	SessionId string `json:"sessionId"`
}

type SendChatResponse struct {
	Text      string `json:"text"`
	Timestamp int64  `json:"timestamp"`
	PlayerId  string `json:"playerId"`
}

type LoadChatResponse struct {
	Messages []SendChatResponse `json:"messages"`
}

type JoinResponse struct {
	LobbyId string `json:"lobbyId"`
}

type GetResponse struct {
	Lobby  Details `json:"lobby"`
	Player Player  `json:"player"`
}

type PlayerJoinResponse struct {
	Player Player `json:"player"`
}

type PlayerLeftResponse struct {
	Player Player `json:"player"`
}
