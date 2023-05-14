package player

type Player struct {
	LobbyId      string `dynamodbav:"LobbyId" json:"-"`
	SessionId    string `dynamodbav:"SessionId" json:"-"`
	ConnectionId string `dynamodbav:"ConnectionId" json:"-"`
	Id           string `dynamodbav:"Id" json:"id"`
	Name         string `dynamodbav:"Name" json:"name"`
	Points       int32  `dynamodbav:"Points" json:"points"`
	IsAdmin      bool   `dynamodbav:"IsAdmin" json:"isAdmin"`
	IsOnline     bool   `dynamodbav:"IsOnline" json:"isOnline"`
}
