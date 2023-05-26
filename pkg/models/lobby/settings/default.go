package settings

import (
	"backend/pkg/models/game"
	"encoding/json"
)

type Encoded string

type BaseSettings struct {
	GameId     game.Id `json:"gameId"`
	MaxPlayers int     `json:"maxPlayers"`
}

func GetDefaultSettings(maxPlayers int) *BaseSettings {
	return &BaseSettings{
		MaxPlayers: maxPlayers,
		GameId:     game.WordGuess,
	}
}

func (s BaseSettings) Encode() (Encoded, error) {
	temp, err := json.Marshal(s)
	return Encoded(temp), err
}

func (str *Encoded) GetGameId() (game.Id, error) {
	settings := BaseSettings{}
	err := json.Unmarshal([]byte(*str), &settings)
	if err != nil {
		return 0, err
	}
	return settings.GameId, nil
}

func (str *Encoded) String() string {
	return string(*str)
}
