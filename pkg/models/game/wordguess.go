package game

import (
	"backend/pkg/models/settings"
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
)

type CardColour string

const (
	Red    CardColour = "red"
	Blue   CardColour = "blue"
	Black  CardColour = "black"
	White  CardColour = "white"
	Hidden CardColour = ""
)

type CardPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type WordGuessState struct {
	Cards   [][]*Card `json:"cards"`
	XLength int       `json:"xLength"`
	YLength int       `json:"yLength"`
}

type Card struct {
	Word     string     `json:"word"`
	Colour   CardColour `json:"colour"`
	Revealed bool       `json:"revealed"`
}

type cardGenerator struct {
	settings      *settings.WordGuessSettings
	cardCount     map[CardColour]int
	blackCards    int
	whiteCards    int
	colouredCards int
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
	xLen := getBoardWidth(settings.TotalCards())
	yLen := settings.TotalCards() / xLen
	state := WordGuessState{
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

func (cg *cardGenerator) Generate(yLen int, xLen int) [][]*Card {
	words, err := loadWords()
	if err != nil {
		log.Printf("Error loading words: %v", err)
	}
	Cards := make([][]*Card, yLen)
	for y := 0; y < yLen; y++ {
		Cards[y] = make([]*Card, xLen)
		for x := 0; x < xLen; x++ {
			Cards[y][x] = &Card{
				Colour:   cg.getColour(),
				Revealed: false,
				Word:     words[rand.Intn(len(words))],
			}
		}
	}
	return Cards
}

func (cg *cardGenerator) getColour() CardColour {

	var cardColours []CardColour
	if cg.blackCards != cg.cardCount[Black] {
		cardColours = append(cardColours, Black)
	}
	if cg.whiteCards != cg.cardCount[White] {
		cardColours = append(cardColours, White)
	}
	if cg.colouredCards != cg.cardCount[Red] {
		cardColours = append(cardColours, Red)
	}
	if cg.colouredCards != cg.cardCount[Blue] {
		cardColours = append(cardColours, Blue)
	}

	if len(cardColours) == 0 {
		//this should never happen, but just in case I did something stupid
		panic("Attempting to generate when no more cards are available")
	}

	colour := cardColours[rand.Intn(len(cardColours))]
	cg.cardCount[colour]++
	return colour
}

func (w *WordGuessState) HideCardColours() *WordGuessState {
	//iterate over cards and hide colours
	for y := 0; y < len(w.Cards); y++ {
		for x := 0; x < len(w.Cards[y]); x++ {
			if !w.Cards[y][x].Revealed {
				w.Cards[y][x].Colour = Hidden
			}
		}
	}
	return w
}
