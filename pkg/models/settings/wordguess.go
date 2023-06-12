package settings

import (
	"backend/pkg/models"
	"encoding/json"
)

type Settings models.Settings

type WordGuessSettings struct {
	BlackCards int `json:"BlackCards"`
	WhiteCards int `json:"WhiteCards"`
}

func GetDefaultWordGuess() *models.Settings {
	settings := models.GetDefaultSettings(12, models.WordGuess)
	settings.Teams = 2
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
