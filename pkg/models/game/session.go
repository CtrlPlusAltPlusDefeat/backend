package game

func (g *Session) IncrementState() *Session {
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

			if currentTurnIndex == len(g.Teams)-1 {
				g.State.CurrentTurn = g.Teams[0].Name
			} else {
				g.State.CurrentTurn = g.Teams[currentTurnIndex+1].Name
			}

			break
		}

	}

	return g
}
