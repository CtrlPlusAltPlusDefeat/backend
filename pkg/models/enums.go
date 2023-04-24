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

// //// Lobby
type lobbyClientAction struct {
	Create string
	Join   string
}
type lobbyServerAction struct {
	Joined       string
	PlayerJoined string
	PlayerLeft   string
}

type lobbyModel struct {
	ClientActions lobbyClientAction
	ServerActions lobbyServerAction
}

var Lobby = lobbyModel{
	ClientActions: lobbyClientAction{
		Create: "create",
		Join:   "join",
	},
	ServerActions: lobbyServerAction{
		Joined:       "joined",
		PlayerJoined: "player-joined",
		PlayerLeft:   "player-left",
	},
}
