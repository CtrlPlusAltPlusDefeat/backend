package state

import (
	"backend/pkg/models"
	"fmt"
)

func ValidateGameState(next ControllerHandler) ControllerHandler {
	return func(ctx *models.Context, controller *Controller) {
		fmt.Printf("Validating GameState\n")
		if controller.GameState != nil {
			next(ctx, controller)
		}
		fmt.Printf("GameState is nil\n")
	}
}
