package chat

type clientAction struct {
	Send string
}
type serverAction struct {
	Receive string
}

type chatAction struct {
	Client clientAction
	Server serverAction
}

var Actions = chatAction{
	Client: clientAction{
		Send: "send",
	},
	Server: serverAction{
		Receive: "receive",
	},
}
