package lobby

type Player struct {
	LobbyId      string `dynamodbav:"LobbyId"`
	SessionId    string `dynamodbav:"SessionId"`
	ConnectionId string `dynamodbav:"ConnectionId"`
	Id           string `dynamodbav:"Id" json:"id"`
	Name         string `dynamodbav:"Name" json:"name"`
	Points       int32  `dynamodbav:"Points" json:"points"`
	IsAdmin      bool   `dynamodbav:"IsAdmin" json:"isAdmin"`
}

type Details struct {
	Players []Player `json:"players"`
	LobbyId string   `json:"lobbyId"`
}
