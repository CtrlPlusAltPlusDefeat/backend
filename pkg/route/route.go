package route

import (
	"backend/pkg/models"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

type SocketData struct {
	requestContext *events.APIGatewayWebsocketProxyRequestContext
	message        models.Wrapper
}

func Route(context *events.APIGatewayWebsocketProxyRequestContext, body string) {
	var message models.Wrapper
	err := message.Decode([]byte(body))

	if err != nil {
		log.Println("Error decoding message", err)
		return
	}

	log.Println("Route ", message.Service)

	routeMessage := SocketData{context, message}

	switch message.Service {
	case models.Service.Chat:
		chatHandle(&routeMessage)
		break
	case models.Service.Player:
		playerHandle(&routeMessage)
	}

}
