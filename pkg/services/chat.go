package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/chat"
	"backend/pkg/ws"
)

func SendChat(context *models.Context, data *models.Data) error {
	req := chat.MessageRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	sender, err := db.LobbyPlayer.Get(context.LobbyId(), context.SessionId())

	if err != nil {
		return err
	}

	c, err := db.LobbyChat.Add(context.LobbyId(), &sender.Id, &req.Text)

	if err != nil {
		return err
	}

	route := models.NewRoute(&models.Service.Chat, &chat.Actions.Server.Receive)
	err = ws.SendToLobby(context, route, chat.MessageResponse{Text: c.Message, Timestamp: c.Timestamp, PlayerId: c.PlayerId})

	if err != nil {
		return err
	}

	return nil
}
