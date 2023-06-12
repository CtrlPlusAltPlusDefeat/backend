package game

import (
	"fmt"
)

func ValidateGameState(next ControllerHandler) ControllerHandler {
	return func(state *Session, controller *Controller) {
		fmt.Printf("Validating GameState\n")
		if controller.GameState != nil {
			next(state, controller)
		}
		fmt.Printf("GameState is nil\n")
	}
}
