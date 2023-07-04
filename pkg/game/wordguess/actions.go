package wordguess

import "backend/pkg/game"

func HandlePlayerAction(ctx *game.Context) (*game.Session, error) {
	err := ctx.Session().IncrementState(*ctx.Player())
	if err != nil {
		return nil, err
	}

	return &game.Session{
		State: ctx.Session().State,
		Game:  ctx.Session().Game,
	}, nil
}
