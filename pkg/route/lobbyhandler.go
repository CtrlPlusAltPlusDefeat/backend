package route

import (
	"backend/pkg/models"
	"backend/pkg/models/lobby"
	"backend/pkg/services"
	"log"
)

func lobbyHandler(context *models.Context, data *models.Data) {
	log.Printf("lobbyHandler: %s", data.Message.Action)

	var err error

	switch data.Message.Action {
	case lobby.Action.Client.Create:
		err = createLobby(context)
		break
	case lobby.Action.Client.Join:
		err = joinLobby(context, &data.Message)
		break
	case lobby.Action.Client.SetName:
		err = setLobbyName(context, &data.Message)
		break
	case lobby.Action.Client.Get:
		err = getLobby(context, &data.Message)
		break
	}

	if err != nil {
		log.Printf("Error handling lobby request %s", err)
	}
}

func createLobby(context *models.Context) error {
	return services.Create(context)
}

func joinLobby(context *models.Context, message *models.Wrapper) error {
	req := lobby.JoinRequest{}
	err := req.Decode(message)

	if err != nil {
		return err
	}

	context.LobbyId = &req.LobbyId

	return services.Join(context, false)
}

func setLobbyName(context *models.Context, message *models.Wrapper) error {
	req := lobby.SetNameRequest{}
	err := req.Decode(message)

	if err != nil {
		return err
	}

	context.LobbyId = &req.LobbyId

	return services.NameChange(context, &req.Text)
}

func getLobby(context *models.Context, message *models.Wrapper) error {
	req := lobby.GetRequest{}
	err := req.Decode(message)

	if err != nil {
		return err
	}

	context.LobbyId = &req.LobbyId

	return services.Get(context)
}
