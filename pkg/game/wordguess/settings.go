package wordguess

import (
	"backend/pkg/models"
	"encoding/json"
)

type Settings struct {
	BlackCards    int `json:"blackCards"`
	WhiteCards    int `json:"whiteCards"`
	ColouredCards int `json:"colouredCards"`
}

func GetDefaultSettings() *models.Settings {
	settings := models.GetDefaultSettings(12, models.WordGuess)
	settings.Teams = 2
	settings.Game, _ = json.Marshal(Settings{
		//black card ends the game instantly
		BlackCards: 1,
		//white card don't give scored
		WhiteCards: 15,
		//when a team has revealed all their cards they win
		ColouredCards: 7,
	})
	return settings
}

func GetSettings(s *models.Settings) (*Settings, error) {
	settings := Settings{}
	err := json.Unmarshal(s.Game, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

func (w *Settings) TotalCards() int {
	//2 times and each team must have the same number of cards
	return w.BlackCards + w.WhiteCards + (w.ColouredCards * 2)
}
