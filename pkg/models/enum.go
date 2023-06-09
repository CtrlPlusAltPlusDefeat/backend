package models

type service struct {
	Player string
	Chat   string
	Lobby  string
	Game   string
}

var Service = service{
	Player: "player",
	Chat:   "chat",
	Lobby:  "lobby",
	Game:   "game",
}
