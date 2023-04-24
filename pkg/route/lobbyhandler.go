package route

import (
	"backend/pkg/models"
	"backend/pkg/services"
	"log"
)

func lobbyHandler(socketData *models.SocketData) {
	log.Printf("lobbyHandler: %s", socketData.Message.Action)

	var err error
	//inject into services
	services.SocketData = socketData
	switch socketData.Message.Action {
	case models.Lobby.ClientActions.Create:
		err = createLobby()
		break
	case models.Lobby.ClientActions.Join:
		err = joinLobby(&socketData.Message)
		break
	case models.Lobby.ClientActions.SetName:
		err = setLobbyName(&socketData.Message)
		break
	case models.Lobby.ClientActions.Get:
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
	req := models.LobbyJoinRequest{}
	err := req.Decode(message)
	if err != nil {
		return err
	}
	return services.Lobby.Join(req.LobbyId, false)
}

func setLobbyName(message *models.Wrapper) error {
	req := models.LobbySetNameRequest{}
	err := req.Decode(message)
	if err != nil {
		return err
	}
	return services.Lobby.NameChange(&req.Text, &req.LobbyId)
}

func getLobby(message *models.Wrapper) error {
	req := models.LobbyGetRequest{}
	err := req.Decode(message)
	if err != nil {
		return err
	}
	return services.Lobby.Get(&req.LobbyId)
}
