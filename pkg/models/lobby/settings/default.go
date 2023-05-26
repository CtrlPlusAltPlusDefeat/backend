package settings

import (
	"backend/pkg/models/game"
	"encoding/json"
)

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

func (s BaseSettings) Encode() ([]byte, error) {
	return json.Marshal(s)
}

type LobbySettings interface {
	BaseSettings | WordGuessSettings
}

func GetType(str string) (game.Id, error) {
	settings := BaseSettings{}
	err := json.Unmarshal([]byte(str), &settings)
	if err != nil {
		return 0, err
	}
	return settings.GameId, nil
}
