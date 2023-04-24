package route

import (
	"backend/pkg/models"
	"backend/pkg/services"
	"backend/pkg/ws"
	"context"
	"github.com/google/uuid"
	"log"
)

func playerHandle(socketData *SocketData) {
	log.Printf("playerHandle: %s", socketData.message.Action)
	var err error
	switch socketData.message.Action {
	case models.Player.ClientActions.CreateSession:
		err = createSession(socketData)
		break
	case models.Player.ClientActions.UseSession:
		err = useSession(socketData)
		break
	}
	if err != nil {
		errorRes, err := models.ErrorResponse{Error: "Something went wrong handling this"}.UseWrapper(socketData.message)
		err = ws.Send(context.TODO(), &socketData.requestContext.ConnectionID, errorRes)
		log.Print(err)
	}
}

func createSession(socketData *SocketData) error {
	return services.Player.CreateSession(socketData.requestContext.ConnectionID)
}

func useSession(socketData *SocketData) error {
	useSessionReq := models.SessionUseRequest{}
	err := useSessionReq.Decode(&socketData.message)
	if err != nil {
		return err
	}
	_, err = uuid.Parse(useSessionReq.SessionId)
	if err != nil {
		return err
	}
	return services.Player.SetSession(useSessionReq.SessionId, socketData.requestContext.ConnectionID)

}
