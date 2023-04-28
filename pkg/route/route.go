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
		return
	}

	res, err := db.Connection.Get(context.ConnectionId())

	if err != nil {
		return
	}

	switch data.Message.Service {
	case models.Service.Chat:
		chatHandle(context.ForSession(&res.SessionId), data)
		break
	case models.Service.Lobby:
		lobbyHandler(context.ForSession(&res.SessionId), data)
		break
	}
}
