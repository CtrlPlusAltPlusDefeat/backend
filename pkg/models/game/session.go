package game

func (g *Session) SetNextTurn() *Session {
	currentTurnIndex := g.Teams.GetIndex(g.State.CurrentTurn)

	if currentTurnIndex == -1 || currentTurnIndex == len(g.Teams)-1 {
		g.State.CurrentTurn = g.Teams[0].Name
	} else {
		g.State.CurrentTurn = g.Teams[currentTurnIndex+1].Name
	}
	return g
}
