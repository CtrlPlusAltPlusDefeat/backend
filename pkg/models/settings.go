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

func (s *Settings) Encode() (Encoded, error) {
	str, err := json.Marshal(s)
	return Encoded(str), err
}

func (s *Settings) DecodeTo(i any) error {
	return json.Unmarshal(s.Game, i)
}
