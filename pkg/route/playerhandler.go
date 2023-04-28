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
	useSessionReq := player.SessionUseRequest{}
	err := useSessionReq.Decode(&data.Message)

	if err != nil {
		return err
	}

	_, err = uuid.Parse(useSessionReq.SessionId)

	if err != nil {
		return err
	}

	return services.SetSession(context.ForSession(&useSessionReq.SessionId))
}
