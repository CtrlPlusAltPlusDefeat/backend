package game

type Session struct {
	LobbyId       string `dynamodbav:"LobbyId" json:"-"`
	GameSessionId string `dynamodbav:"GameSessionId" json:"-"`
	GameTypeId    Id     `dynamodbav:"ConnectionId" json:"-"`
}

type Team struct {
	Name    string
	Players []string
}

type TeamSession struct {
	Session
	Teams []Team
}
