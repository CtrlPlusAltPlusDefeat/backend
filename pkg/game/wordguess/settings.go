package wordguess

import (
	"backend/pkg/models"
	"encoding/json"
)

func GetDefaultSettings() *models.Settings {
	settings := models.GetDefaultSettings(12, models.WordGuess)
	settings.Teams = 2
	settings.Game, _ = (&models.WordGuessSettings{
		//black card ends the game instantly
		BlackCards: 1,
		//white card don't give scored
		WhiteCards: 15,
		//when a team has revealed all their cards they win
		ColouredCards: 7,
	}).Encode()
	return settings
}

func GetSettings(s *models.Settings) (*models.WordGuessSettings, error) {
	settings := models.WordGuessSettings{}
	err := json.Unmarshal(s.Game, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}
