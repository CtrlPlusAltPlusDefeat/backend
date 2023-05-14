package chat

type Chat struct {
	LobbyId   string `dynamodbav:"LobbyId" json:"lobbyId"`
	PlayerId  string `dynamodbav:"PlayerId" json:"playerId"`
	Timestamp string `dynamodbav:"Timestamp" json:"timestamp"`
	Message   string `dynamodbav:"Message" json:"message"`
}
