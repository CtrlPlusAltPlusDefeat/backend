package chat

type Chat struct {
	LobbyId   string `dynamodbav:"LobbyId" json:"lobbyId"`
	PlayerId  string `dynamodbav:"PlayerId" json:"playerId"`
	Timestamp int64  `dynamodbav:"Timestamp" json:"timestamp"`
	Message   string `dynamodbav:"Message" json:"message"`
}
