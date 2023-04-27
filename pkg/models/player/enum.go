package player

type clientAction struct {
	CreateSession string
	UseSession    string
}
type serverActions struct {
	SetSession string
}

type playerAction struct {
	Client clientAction
	Server serverActions
}

var Action = playerAction{
	Client: clientAction{
		CreateSession: "create-session",
		UseSession:    "use-session",
	},
	Server: serverActions{
		SetSession: "set-session",
	},
}
