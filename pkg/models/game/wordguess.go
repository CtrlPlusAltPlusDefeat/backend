package game

import (
	"backend/pkg/models/settings"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
)

type CardPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type WordGuessState struct {
	Cards         [][]*Card       `json:"cards"`
	RevealedCards []*CardPosition `json:"revealedCards"`
}

type Card struct {
	Word   string `json:"word"`
	Colour string `json:"colour"`
	Type   string `json:"type"`
}

func (w *WordGuessState) Encode() ([]byte, error) {
	return json.Marshal(w)
}

func (s *Session) GetWordGuess() (*WordGuessState, error) {
	state := WordGuessState{}
	err := json.Unmarshal(s.Game, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func NewWordGuessState(settings *settings.WordGuessSettings) *WordGuessState {
	state := WordGuessState{
		Cards:         make([][]*Card, settings.TotalCards()),
		RevealedCards: make([]*CardPosition, 0),
	}
	state.GenerateCards(settings)
	return &state
}

func getBoardWidth(totalCards int) int {
	// don't have more than 7 and less than 3 cards in a row
	maxX := 7
	minX := 3
	factor := minX
	//get factors between 3 and 7
	for i := minX; i <= maxX; i++ {
		if totalCards%i == 0 {
			factor = i
		}
	}

	return factor
}

func loadWords() ([]string, error) {
	wd, _ := os.Getwd()
	var filepath string
	if os.Getenv("LOCAL_WEBSOCKET_SERVER") == "1" {
		filepath = wd + "/wordpacks/pack1.txt"
	} else {
		filepath = wd + "./wordpacks/pack1.txt"
	}

	contents, err := os.ReadFile(filepath)
	log.Printf("Loading words from %v", contents)
	if err != nil {
		return nil, err
	}

	words := strings.Split(strings.ReplaceAll(string(contents), "\r\n", "\n"), "\n")
	return words, nil
}

func (w *WordGuessState) GenerateCards(settings *settings.WordGuessSettings) *WordGuessState {
	xLen := getBoardWidth(settings.TotalCards())
	yLen := settings.TotalCards() / xLen
	words, err := loadWords()
	if err != nil {
		log.Printf("Error loading words: %v", err)
	}
	w.Cards = make([][]*Card, yLen+1)
	for y := 0; y <= yLen; y++ {
		w.Cards[y] = make([]*Card, xLen+1)
		for x := 0; x <= xLen; x++ {
			w.AddCard(&CardPosition{X: x, Y: y}, &Card{
				Colour: "",
				Type:   "",
				Word:   words[rand.Intn(len(words))],
			})
		}
	}
	return w
}

func (w *WordGuessState) AddCard(pos *CardPosition, card *Card) *WordGuessState {
	w.Cards[pos.Y][pos.X] = card
	return w
}
