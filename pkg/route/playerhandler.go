package route

import (
	"backend/pkg/models"
	"backend/pkg/services"
	"log"
)

func playerHandle(socketData *SocketData) {
	var err error
	switch socketData.message.Action {
	case models.Player.ClientActions.CreateSession:
		createSession(socketData)
		break
	case models.Player.ClientActions.UseSession:
		err = useSession(socketData)
		break
	}
	if err != nil {
		log.Print(err)
	}
}

func createSession(socketData *SocketData) {
	services.Player.CreateSession(socketData.requestContext.ConnectionID)
}

func useSession(socketData *SocketData) error {
	useSessionReq := models.SessionUseRequest{}
	err := useSessionReq.Decode(&socketData.message)
	if err != nil {
		return err
	}
	services.Player.SetSession(useSessionReq.SessionId, socketData.requestContext.ConnectionID)
	return nil
}
