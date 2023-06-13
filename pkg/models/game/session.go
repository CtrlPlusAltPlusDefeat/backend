package game

import (
	"backend/pkg/models"
	"fmt"
)

func (s *Session) IncrementState(player models.Player) error {
	switch s.State.State {
	case models.PreMatch:
		{
			if !player.IsAdmin {
				return fmt.Errorf("player is not an admin")
			}

			s.State.CurrentTurn = s.Teams.GetRandom().Name
			s.State.State = models.InProgress
			break
		}
	case models.InProgress:
		{
			currentTurnIndex := s.Teams.GetIndex(s.State.CurrentTurn)

			if s.Teams[currentTurnIndex].IncludesPlayer(player.Id) == false {
				return fmt.Errorf("player not apart of current team")
			}

			if currentTurnIndex == len(s.Teams)-1 {
				s.State.CurrentTurn = s.Teams[0].Name
			} else {
				s.State.CurrentTurn = s.Teams[currentTurnIndex+1].Name
			}

			break
		}

	}

	return nil
}
