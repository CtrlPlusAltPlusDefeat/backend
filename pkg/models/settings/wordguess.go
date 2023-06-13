package settings

import (
	"backend/pkg/models"
	"encoding/json"
)

type WordGuessSettings struct {
	BlackCards int `json:"BlackCards"`
	WhiteCards int `json:"WhiteCards"`
}

func GetDefaultWordGuess() *models.Settings {
	settings := models.GetDefaultSettings(12, models.WordGuess)
	settings.Teams = 2
	settings.Game, _ = json.Marshal(WordGuessSettings{
		BlackCards: 5,
		WhiteCards: 15,
	})
	return settings
}

func GetWordGuess(s *models.Settings) (*WordGuessSettings, error) {
	settings := WordGuessSettings{}
	err := json.Unmarshal(s.Game, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

func (w *WordGuessSettings) TotalCards() int {
	return w.BlackCards + w.WhiteCards
}
