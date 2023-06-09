package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/ws"
)

func SendChat(context *models.Context, data *models.Data) error {
	req := models.SendChatRequest{}
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

	err = ws.SendToLobby(context, context.Route(), models.SendChatResponse{Text: c.Message, Timestamp: c.Timestamp, PlayerId: c.PlayerId})

	if err != nil {
		return err
	}

	return nil
}

func LoadChat(context *models.Context, data *models.Data) error {
	req := models.LoadChatRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	c, err := db.LobbyChat.Get(context.LobbyId(), req.Timestamp)

	if err != nil {
		return err
	}

	response := models.LoadChatResponse{}

	for _, item := range c {
		response.Messages = append(response.Messages, models.SendChatResponse{Text: item.Message, Timestamp: item.Timestamp, PlayerId: item.PlayerId})
	}
	err = ws.Send(context, context.Route(), response)
	if err != nil {
		return err
	}

	return nil
}
