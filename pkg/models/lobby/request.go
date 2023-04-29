package lobby

type JoinRequest struct {
	LobbyId string `json:"lobbyId"`
}

type GetRequest struct {
	LobbyId string `json:"lobbyId"`
}

type SetNameRequest struct {
	Text    string `json:"text"`
	LobbyId string `json:"lobbyId"`
}
