package models

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

type TeamName string

const (
	None     TeamName = ""
	RedTeam  TeamName = "red"
	BlueTeam TeamName = "blue"
)

func (t TeamName) IsValid() bool {
	if t == RedTeam || t == BlueTeam {
		return true
	}
	return false
}

func GetTeamName(teamNum int) TeamName {
	switch teamNum {
	case 1:
		return BlueTeam
	default:
		return RedTeam
	}
}

type State string

const (
	PreMatch   State = "prematch"
	InProgress State = "inprogress"
	PostMatch  State = "postmatch"
)
