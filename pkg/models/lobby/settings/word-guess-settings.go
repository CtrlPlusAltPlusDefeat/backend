package settings

import (
	"encoding/json"
)

type WordGuessSettings struct {
	BaseSettings
}

func NewWordGuess(str string) (*WordGuessSettings, error) {
	settings := WordGuessSettings{}
	err := json.Unmarshal([]byte(str), &settings)

	if err != nil {
		return nil, err
	}

	return &settings, nil
}
