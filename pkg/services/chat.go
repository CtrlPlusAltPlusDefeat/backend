package services

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/models/chat"
	"backend/pkg/ws"
	"log"
)

func SendChat(context *models.Context, data *models.Data) error {
	req := chat.MessageRequest{}
	err := data.DecodeTo(&req)

	if err != nil {
		return err
	}

	players, err := db.LobbyPlayer.GetPlayers(context.LobbyId())

	if err != nil {
		return err
	}

	route := models.NewRoute(&models.Service.Chat, &chat.Actions.Server.Receive)

	log.Println("Sending ", req.Text, " to ", len(players), " players")

	for index, player := range players {
		log.Println("Sending ", req.Text, " to player ", index)

		err := ws.Send(context.ForConnection(&player.ConnectionId), route, chat.MessageResponse{Text: req.Text, ConnectionId: player.Id})

		if err != nil {
			log.Printf("Error sending: %s", err)
		}
	}

	return nil
}
