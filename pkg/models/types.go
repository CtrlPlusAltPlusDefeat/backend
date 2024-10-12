package models

import "encoding/json"

type Connection struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
	SessionId    string `dynamodbav:"SessionId"`
}

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

type Chat struct {
	LobbyId   string `dynamodbav:"LobbyId" json:"lobbyId"`
	PlayerId  string `dynamodbav:"PlayerId" json:"playerId"`
	Timestamp int64  `dynamodbav:"Timestamp" json:"timestamp"`
	Message   string `dynamodbav:"Message" json:"message"`
}

type Details struct {
	Players       []Player `json:"players"`
	LobbyId       string   `json:"lobbyId"`
	Settings      Settings `json:"settings"`
	InGame        bool     `json:"inGame"`
	GameSessionId string   `json:"gameSessionId"`
}

type Lobby struct {
	LobbyId       string   `json:"lobbyId" dynamodbav:"LobbyId"`
	Settings      Settings `json:"settings" dynamodbav:"Settings"`
	InGame        bool     `json:"inGame" dynamodbav:"InGame"`
	GameSessionId string   `json:"gameSessionId" dynamodbav:"GameSessionId"`
}

type Settings struct {
	GameId     Id              `json:"gameId"`
	MaxPlayers int             `json:"maxPlayers"`
	Teams      int             `json:"teams"`
	Game       json.RawMessage `json:"game" tstype:"WordGuessSettings"`
}
