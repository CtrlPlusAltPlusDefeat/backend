package wordguess

import (
	"backend/pkg/game"
	"backend/pkg/models"
	"backend/pkg/models/context"
	"fmt"
	"log"
)

var HideCardsHandler = context.BeforeSend(func(context *context.Context, player *models.Player, message any) (any, error) {
	// check message is of type game.Session return if not
	gState, ok := message.(*game.Session)
	if !ok {
		return message, fmt.Errorf("message is not of type game.Session")
	}

	log.Printf("Got game state: %v", gState)

	state, err := GetState(gState)
	if err != nil {
		return message, fmt.Errorf("error getting state: %v", err)
	}

	// will need to get player role
	var playerData PlayerData
	err = gState.Teams.GetPlayer(player.Id).DecodeTo(&playerData)
	if err != nil {
		return message, fmt.Errorf("error getting state: %v", err)
	}

	// this is stored on game session using playerlobbyid
	if playerData.Role == SpyMaster && gState.State.State != models.PreMatch {
		return message, nil
	}

	//hide cards from player
	state.HideCardColours()
	encode, err := state.Encode()
	if err != nil {
		return message, err
	}
	gState.Game = encode

	return gState, nil
})

func HandlePlayerAction(ctx *context.Context, data *models.Data, player *models.Player) (*game.Session, error) {
	err := ctx.GameSession().IncrementState(*player)
	if err != nil {
		return nil, err
	}

	//	gState, err := GetState(ctx.GameSession())

	ctx.ForBeforeSend(&HideCardsHandler)

	if err != nil {
		return nil, err
	}

	return &game.Session{
		State: ctx.GameSession().State,
		Game:  ctx.GameSession().Game,
	}, nil
}

func HandleSwapTeam(ctx *context.Context, data *models.Data, player *models.Player) (game.TeamArray, error) {
	var playerData PlayerData
	req := SwapTeamRequest{}
	err := data.DecodeTo(&req)
	if err != nil {
		return nil, err
	}
	teams := ctx.GameSession().Teams.SwapTeam(player.Id, req.Team)

	tIndex := teams.GetIndex(req.Team)
	pIndex := teams[tIndex].GetPlayerIndex(player.Id)

	err = teams[tIndex].Players[pIndex].DecodeTo(&playerData)
	if err != nil {
		return nil, err
	}

	playerData.Role = req.Role
	encodedData, err := playerData.Encode()
	if err != nil {
		return nil, err
	}

	teams[tIndex].Players[pIndex].Data = encodedData
	return teams, nil
}

func HandleGetState(ctx *context.Context, data *models.Data) (*game.Session, error) {
	ctx.ForBeforeSend(&HideCardsHandler)
	return ctx.GameSession(), nil
}
