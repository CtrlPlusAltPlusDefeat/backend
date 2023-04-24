package route

import (
	"backend/pkg/models"
	"backend/pkg/services"
	"log"
)

func lobbyHandler(socketData *SocketData) {

	var err error
	switch socketData.message.Action {
	case models.Lobby.ClientActions.Create:
		err = createLobby(socketData)
		break
	case models.Lobby.ClientActions.Join:
		err = joinLobby(socketData)
		break
	}
	if err != nil {
		log.Printf("Error handling lobby request %s", err)
	}
}

func createLobby(socketData *SocketData) error {
	return services.Lobby.Create(&socketData.requestContext.ConnectionID, socketData.sessionId)
}

func joinLobby(socketData *SocketData) error {
	req := models.LobbyJoinRequest{}
	err := req.Decode(&socketData.message)
	if err != nil {
		return err
	}
	return services.Lobby.Join(req.LobbyId, &socketData.requestContext.ConnectionID, socketData.sessionId, false)
}
