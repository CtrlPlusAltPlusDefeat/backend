package ws

import (
	"backend/pkg/services"
	sockethelpers "backend/pkg/socket-helpers"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

// Message Generic wrapper for all websocket messages

func Route(context *events.APIGatewayWebsocketProxyRequestContext, body string) {
	var message sockethelpers.Message
	err := message.Decode([]byte(body))
	if err != nil {
		fmt.Println("Error decoding message", err)
		return
	}

	fmt.Println("Route ", message.Service)
	switch message.Service {
	case "chat":
		services.HandleChat(context, message)
		break
	}

}
