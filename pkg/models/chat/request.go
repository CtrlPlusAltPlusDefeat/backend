package chat

type MessageRequest struct {
	Text    string `json:"text"`
	LobbyId string `json:"lobbyId"`
}
