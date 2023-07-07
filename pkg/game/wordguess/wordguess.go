package wordguess

import (
	"backend/pkg/game"
	"backend/pkg/models"
	"encoding/json"
)

type State struct {
	Cards   [][]*Card `json:"cards"`
	XLength int       `json:"xLength"`
	YLength int       `json:"yLength"`
}

func (w *State) Encode() ([]byte, error) {
	return json.Marshal(w)
}

func Setup(teams game.TeamArray, settings *models.Settings) (game.TeamArray, []byte, error) {
	wordGuessSettings, err := GetSettings(settings)
	if err != nil {
		return nil, nil, err
	}
	gState, err := newState(wordGuessSettings).Encode()
	if err != nil {
		return nil, nil, err
	}

	teams, err = AddRoleDefaults(teams)
	if err != nil {
		return nil, nil, err
	}

	return teams, gState, nil
}

func GetState(s *game.Session) (*State, error) {
	state := State{}
	err := json.Unmarshal(s.Game, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (w *State) HideCardColours() *State {
	for y := 0; y < len(w.Cards); y++ {
		for x := 0; x < len(w.Cards[y]); x++ {
			if !w.Cards[y][x].Revealed {
				w.Cards[y][x].Colour = Hidden
			}
		}
	}
	return w
}

func newState(settings *Settings) *State {
	xLen := getBoardWidth(settings.TotalCards())
	yLen := settings.TotalCards() / xLen
	state := State{
		XLength: xLen,
		YLength: yLen,
		Cards:   make([][]*Card, settings.TotalCards()),
	}

	generateCards := &cardGenerator{
		settings:      settings,
		cardCount:     make(map[CardColour]int, 0),
		blackCards:    settings.BlackCards,
		whiteCards:    settings.WhiteCards,
		colouredCards: settings.ColouredCards,
	}

	state.Cards = generateCards.Generate(xLen, yLen)
	return &state
}
