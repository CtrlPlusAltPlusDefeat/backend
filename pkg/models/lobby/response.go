package lobby

import "backend/pkg/models/player"

type JoinResponse struct {
	LobbyId string `json:"lobbyId"`
}

type GetResponse struct {
	Lobby  Details       `json:"lobby"`
	Player player.Player `json:"player"`
}

type PlayerJoinResponse struct {
	Player player.Player `json:"player"`
}

type PlayerLeftResponse struct {
	Player player.Player `json:"player"`
}
