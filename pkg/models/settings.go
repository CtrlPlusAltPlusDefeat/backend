package models

import (
	"encoding/json"
)

type Encoded string

func GetDefaultSettings(maxPlayers int, gameId Id) *Settings {
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
