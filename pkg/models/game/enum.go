package game

type Id int

const (
	WordGuess Id = iota
	maxSize      // Always keep this at the end
)

var gameIdToString = map[Id]string{
	WordGuess: "word-guess",
}

func (g Id) String() string {
	return gameIdToString[g]
}

func (g Id) Valid() bool {
	return g > 0 && g < maxSize
}
