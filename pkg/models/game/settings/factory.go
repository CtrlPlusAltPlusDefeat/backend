package settings

import (
	"backend/pkg/models/game"
	"encoding/json"
)

type Encoded string

type Settings struct {
	GameId     game.Id         `json:"gameId"`
	MaxPlayers int             `json:"maxPlayers"`
	Teams      int             `json:"teams"`
	Other      json.RawMessage `json:"other"`
}

type BaseSettings struct {
	GameId     game.Id `json:"gameId"`
	MaxPlayers int     `json:"maxPlayers"`
	Teams      int     `json:"teams"`
}

func GetDefaultSettings(maxPlayers int, gameId game.Id) *Settings {
	return &Settings{
		MaxPlayers: maxPlayers,
		GameId:     gameId,
	}
}

func (str *Encoded) Decode() (*Settings, error) {
	settings := Settings{}
	err := json.Unmarshal([]byte(*str), &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

func (s *Settings) Encode() ([]byte, error) {
	return json.Marshal(s)
}

func (s *Settings) GetBaseSettings() *BaseSettings {
	return &BaseSettings{
		GameId:     s.GameId,
		MaxPlayers: s.MaxPlayers,
		Teams:      s.Teams,
	}
}
