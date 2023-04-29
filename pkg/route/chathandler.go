package route

import (
	"backend/pkg/models"
	"backend/pkg/models/chat"
	"backend/pkg/services"
)

func SendChat(context *models.Context, data *models.Data) error {
	req := chat.MessageRequest{}
	err := data.DecodeTo(req)

	if err != nil {
		return err
	}

	return services.BroadcastMessage(context, req)
}
