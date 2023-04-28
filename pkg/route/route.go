package route

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"log"
)

func Route(context *models.Context, data *models.Data) {

	log.Println("Route ", data.Message.Service)

	if data.Message.Service == models.Service.Player {
		playerHandle(context, data)
	}

	res, err := db.Connection.Get(context.Connection.Id)

	if err != nil {
		return
	}

	context.SessionId = &res.SessionId

	switch data.Message.Service {
	case models.Service.Chat:
		chatHandle(context, data)
		break
	case models.Service.Lobby:
		lobbyHandler(context, data)
		break
	}
}
