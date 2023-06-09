package settings

import (
	"backend/pkg/models/game"
	"encoding/json"
)

type WordGuessSettings struct {
	BlackCards int `json:"BlackCards"`
	WhiteCards int `json:"WhiteCards"`
}

func GetDefaultWordGuess() *Settings {
	settings := GetDefaultSettings(12, game.WordGuess)
	settings.Other, _ = json.Marshal(WordGuessSettings{
		BlackCards: 10,
		WhiteCards: 10,
	})
	return settings
}

func (s *Settings) GetWordGuess() (*WordGuessSettings, error) {
	settings := WordGuessSettings{}
	err := json.Unmarshal(s.Other, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}
