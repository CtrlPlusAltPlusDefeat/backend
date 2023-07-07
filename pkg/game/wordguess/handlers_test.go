package wordguess

import (
	"backend/pkg/game"
	"backend/pkg/models"
	ctx "backend/pkg/models/context"
	"context"
	"encoding/json"
	"testing"
)

var connectionId = ""
var connectionHost = ""
var connectionPath = ""
var handlerCtx = ctx.NewContext(context.TODO(), &connectionId, &connectionHost, &connectionPath)
var teams = []game.Team{
	{
		Name:    models.RedTeam,
		Players: []game.TeamPlayer{{Id: "1", Data: json.RawMessage(`{"role":"operative"}`)}},
	},
	{
		Name:    models.BlueTeam,
		Players: []game.TeamPlayer{{Id: "2", Data: json.RawMessage(`{"role":"spymaster"}`)}},
	}}
var state = &State{
	Cards:   [][]*Card{{{Word: "test", Revealed: false, Colour: Red}}},
	XLength: 1,
	YLength: 1,
}
var gameData, _ = state.Encode()
var gameSession = game.Session{
	Info: &game.SessionInfo{
		GameTypeId: models.WordGuess,
	},
	State: game.NewGameState(),
	Teams: teams,
	Game:  gameData,
}
var player1 = models.Player{
	Id:   "1",
	Name: "Tom",
}
var player2 = models.Player{
	Id:   "2",
	Name: "Jerry",
}

// TestHideCardsHandlerNoValue calls HideCardsHandler with null
func TestHideCardsHandlerNoValue(t *testing.T) {
	handler, err := HideCardsHandler(handlerCtx, nil, nil)
	if err != nil {
		return
	}
	if handler != nil || err != nil {
		t.Logf("HideCardsHandler returned a value when it should not have: %v", handler)
	}
}

// TestHideCardsHandlerNoValue should hide cards in PreMatch state
func TestHideCardsHandlerPreMatch(t *testing.T) {
	//copy the struct as we don't want to modify the original
	var gSession = gameSession
	gSession.State.State = models.PreMatch
	handler, err := HideCardsHandler(handlerCtx, &player1, &gSession)
	if err != nil {
		return
	}
	gState, ok := handler.(*game.Session)
	if !ok {
		t.Logf("incorrect value")
		t.FailNow()
	}
	wState, err := GetState(gState)
	if wState.Cards[0][0].Colour != Hidden {
		t.Logf("incorrect colour")
		t.FailNow()
	}
}

func TestHideCardsHandlerInPlayOperative(t *testing.T) {
	//copy the struct as we don't want to modify the original
	var gSession = gameSession
	gSession.State.State = models.InProgress
	handler, err := HideCardsHandler(handlerCtx, &player1, &gSession)
	if err != nil {
		return
	}
	gState, ok := handler.(*game.Session)
	if !ok {
		t.Logf("incorrect value")
		t.FailNow()
	}
	wState, err := GetState(gState)
	if wState.Cards[0][0].Colour != Hidden {
		t.Logf("incorrect colour")
		t.FailNow()
	}
}
func TestHideCardsHandlerInPlaySpyMaster(t *testing.T) {
	//copy the struct as we don't want to modify the original
	var gSession = gameSession
	gSession.State.State = models.InProgress
	handler, err := HideCardsHandler(handlerCtx, &player2, &gSession)
	if err != nil {
		return
	}
	gState, ok := handler.(*game.Session)
	if !ok {
		t.Logf("incorrect value")
		t.FailNow()
	}
	wState, err := GetState(gState)
	if wState.Cards[0][0].Colour != Red {
		t.Logf("incorrect colour, got %s and was expecting %s", wState.Cards[0][0].Colour, Red)
		t.FailNow()
	}
}
