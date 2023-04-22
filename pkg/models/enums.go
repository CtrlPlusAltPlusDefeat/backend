package models

type service struct {
	Player string
	Chat   string
}

var Service = service{
	Player: "player",
	Chat:   "chat",
}

////// Player

type playerClientActions struct {
	CreateSession string
	UseSession    string
}
type playerServerActions struct {
	SetSession string
}

type playerModel struct {
	ClientActions playerClientActions
	ServerActions playerServerActions
}

var Player = playerModel{
	ClientActions: playerClientActions{
		CreateSession: "create-session",
		UseSession:    "use-session",
	},
	ServerActions: playerServerActions{
		SetSession: "set-session",
	},
}

////// Chat

type chatClientActions struct {
	Send string
}
type chatServerActions struct {
	Receive string
}

type chatModel struct {
	ClientActions chatClientActions
	ServerActions chatServerActions
}

var Chat = chatModel{
	ClientActions: chatClientActions{
		Send: "send",
	},
	ServerActions: chatServerActions{
		Receive: "receive",
	},
}
