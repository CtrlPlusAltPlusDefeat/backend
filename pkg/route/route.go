package route

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/ws"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"log"
)

func Route(context *events.APIGatewayWebsocketProxyRequestContext, body string) {
	var message models.Wrapper
	err := message.Decode([]byte(body))

	if err != nil {
		log.Println("Error decoding message", err)
		return
	}

	log.Println("Route ", message.Service)

	//inject ConnectionContext into context
	ws.ConnectionContext = context

	if message.Service == models.Service.Player {
		playerHandle(&models.SocketData{RequestContext: context, Message: message})
	}

	res, err := db.Connection.Get(context.ConnectionID)
	if err != nil {
		return
	}
	routeMessage := models.SocketData{RequestContext: context, Message: message, SessionId: aws.String(res.SessionId)}
	switch message.Service {
	case models.Service.Chat:
		chatHandle(&routeMessage)
		break
	case models.Service.Lobby:
		lobbyHandler(&routeMessage)
		break
	}

}
