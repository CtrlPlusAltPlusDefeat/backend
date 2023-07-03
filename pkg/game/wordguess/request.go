package wordguess

import "backend/pkg/models"

type SwapTeamRequest struct {
	Team models.TeamName `json:"team"` //the team to change to
	Role role            `json:"role"` //the team to change to
}
