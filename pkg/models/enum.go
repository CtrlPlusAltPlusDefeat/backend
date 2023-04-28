package models

type service struct {
	Player string
	Chat   string
	Lobby  string
}

var Service = service{
	Player: "player",
	Chat:   "chat",
	Lobby:  "lobby",
}
