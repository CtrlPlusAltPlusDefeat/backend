package lobby

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

type NameChangeResponse struct {
	Player Player `json:"player"`
}
