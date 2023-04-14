package service

import (
	"backend/pkg/ws"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func Route(context *events.APIGatewayWebsocketProxyRequestContext, body string) {
	var message ws.Message
	_, err := message.Decode([]byte(body))
	if err != nil {
		return
	}

	fmt.Println("Route ", message.Service)
	switch message.Service {
	case "chat":
		HandleChat(context, message)
		break
	}

}
