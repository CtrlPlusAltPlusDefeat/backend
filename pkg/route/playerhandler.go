package route

import (
	"backend/pkg/models"
	"backend/pkg/models/player"
	"backend/pkg/services"
	"backend/pkg/ws"
	"github.com/google/uuid"
	"log"
)

func playerHandle(context *models.Context, data *models.Data) {
	log.Printf("playerHandle: %s", data.Message.Action)

	var err error

	switch data.Message.Action {
	case player.Action.Client.CreateSession:
		err = createSession(context)
		break
	case player.Action.Client.UseSession:
		err = useSession(context, data)
		break
	}

	if err != nil {
		errorRes, err := models.ErrorResponse{Error: "Something went wrong handling this"}.UseWrapper(data.Message)
		err = ws.Send(context, errorRes)
		log.Print(err)
	}
}

func createSession(context *models.Context) error {
	return services.CreateSession(context)
}

func useSession(context *models.Context, data *models.Data) error {
	useSessionReq := player.SessionUseRequest{}
	err := useSessionReq.Decode(&data.Message)

	if err != nil {
		return err
	}

	_, err = uuid.Parse(useSessionReq.SessionId)

	if err != nil {
		return err
	}

	return services.SetSession(context.ForSession(&useSessionReq.SessionId))
}
