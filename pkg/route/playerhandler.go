package route

import (
	"backend/pkg/models"
	"backend/pkg/models/player"
	"backend/pkg/services"
	"github.com/google/uuid"
)

func CreateSession(context *models.Context, data *models.Data) error {
	return services.CreateSession(context)
}

func UseSession(context *models.Context, data *models.Data) error {
	req := player.SessionUseRequest{}
	err := data.DecodeTo(req)

	if err != nil {
		return err
	}

	_, err = uuid.Parse(req.SessionId)

	if err != nil {
		return err
	}

	return services.SetSession(context.ForSession(&req.SessionId))
}
