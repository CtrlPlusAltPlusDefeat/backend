package game

import (
	"backend/pkg/models"
	"fmt"
)

func (g *Session) IncrementState(player models.Player) error {
	switch g.State.State {
	case "prematch":
		{
			g.State.CurrentTurn = g.Teams[0].Name
			g.State.State = "inprogress"

			break
		}
	case "inprogress":
		{
			currentTurnIndex := g.Teams.GetIndex(g.State.CurrentTurn)

			if g.Teams[currentTurnIndex].IncludesPlayer(player.Id) == false {
				return fmt.Errorf("player not apart of current team")
			}

			if currentTurnIndex == len(g.Teams)-1 {
				g.State.CurrentTurn = g.Teams[0].Name
			} else {
				g.State.CurrentTurn = g.Teams[currentTurnIndex+1].Name
			}

			break
		}

	}

	return nil
}
