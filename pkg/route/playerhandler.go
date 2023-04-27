package route

import (
	"backend/pkg/models"
	"backend/pkg/models/player"
	"backend/pkg/services"
	"backend/pkg/ws"
	"context"
	"github.com/google/uuid"
	"log"
)

func playerHandle(socketData *models.SocketData) {
	log.Printf("playerHandle: %s", socketData.Message.Action)
	var err error
	switch socketData.Message.Action {
	case player.Action.Client.CreateSession:
		err = createSession(socketData)
		break
	case player.Action.Client.UseSession:
		err = useSession(socketData)
		break
	}
	if err != nil {
		errorRes, err := models.ErrorResponse{Error: "Something went wrong handling this"}.UseWrapper(socketData.Message)
		err = ws.Send(context.TODO(), &socketData.RequestContext.ConnectionID, errorRes)
		log.Print(err)
	}
}

func createSession(socketData *models.SocketData) error {
	return services.Player.CreateSession(socketData.RequestContext.ConnectionID)
}

func useSession(socketData *models.SocketData) error {
	useSessionReq := player.SessionUseRequest{}
	err := useSessionReq.Decode(&socketData.Message)
	if err != nil {
		return err
	}
	_, err = uuid.Parse(useSessionReq.SessionId)
	if err != nil {
		return err
	}
	return services.Player.SetSession(useSessionReq.SessionId, socketData.RequestContext.ConnectionID)

}
