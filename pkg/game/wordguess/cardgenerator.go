package wordguess

import (
	"log"
	"math/rand"
	"os"
	"strings"
)

type CardColour string

type Words []string

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

type Card struct {
	Word     string     `json:"word"`
	Colour   CardColour `json:"colour"`
	Revealed bool       `json:"revealed"`
}

type cardGenerator struct {
	settings      *Settings
	cardCount     map[CardColour]int
	blackCards    int
	whiteCards    int
	colouredCards int
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

func loadWords() (*Words, error) {
	wd, _ := os.Getwd()
	filepath := wd + "/wordpacks/pack1.txt"

	contents, err := os.ReadFile(filepath)
	log.Printf("Loading words from %v", contents)
	if err != nil {
		return nil, err
	}

	words := Words(strings.Split(strings.ReplaceAll(string(contents), "\r\n", "\n"), "\n"))
	return &words, nil
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
				Word:     words.GetRandom(),
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

func (w *Words) GetRandom() string {
	words := *w
	index := rand.Intn(len(words))
	word := words[index]
	//remove word from array
	*w = append(words[:index], words[index+1:]...)
	return word
}
