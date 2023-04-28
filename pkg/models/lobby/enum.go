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
	Get          string
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
		Get:     "get",
	},
	Server: serverAction{
		Joined:       "joined",
		PlayerJoined: "player-joined",
		PlayerLeft:   "player-left",
		NameChanged:  "name-change",
		Get:          "get",
	},
}