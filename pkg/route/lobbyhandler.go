package route

import (
	"backend/pkg/models"
	"backend/pkg/models/lobby"
	"backend/pkg/services"
	"log"
)

func lobbyHandler(socketData *models.SocketData) {
	log.Printf("lobbyHandler: %s", socketData.Message.Action)

	var err error
	//inject into services
	services.SocketData = socketData
	switch socketData.Message.Action {
	case lobby.Action.Client.Create:
		err = createLobby()
		break
	case lobby.Action.Client.Join:
		err = joinLobby(&socketData.Message)
		break
	case lobby.Action.Client.SetName:
		err = setLobbyName(&socketData.Message)
		break
	case lobby.Action.Client.Get:
		err = getLobby(&socketData.Message)
		break
	}
	if err != nil {
		log.Printf("Error handling lobby request %s", err)
	}
}

func createLobby() error {
	return services.Lobby.Create()
}

func joinLobby(message *models.Wrapper) error {
	req := lobby.JoinRequest{}
	err := req.Decode(message)
	if err != nil {
		return err
	}
	return services.Lobby.Join(req.LobbyId, false)
}

func setLobbyName(message *models.Wrapper) error {
	req := lobby.SetNameRequest{}
	err := req.Decode(message)
	if err != nil {
		return err
	}
	return services.Lobby.NameChange(&req.Text, &req.LobbyId)
}

func getLobby(message *models.Wrapper) error {
	req := lobby.GetRequest{}
	err := req.Decode(message)
	if err != nil {
		return err
	}
	return services.Lobby.Get(&req.LobbyId)
}
