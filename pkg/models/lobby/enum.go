package lobby

type clientAction struct {
	Create  string
	Join    string
	SetName string
	Get     string
}
type serverAction struct {
	Joined       string
	PlayerJoined string
	PlayerLeft   string
	NameChanged  string
	LoadGame     string
}

type lobbyAction struct {
	Client clientAction
	Server serverAction
}

var Action = lobbyAction{
	Client: clientAction{
		Create:  "create",
		Join:    "join",
		SetName: "set-name",
	},
	Server: serverAction{
		Joined:       "join",
		PlayerJoined: "player-joined",
		PlayerLeft:   "player-left",
		NameChanged:  "name-change",
		LoadGame:     "load-game",
	},
}
