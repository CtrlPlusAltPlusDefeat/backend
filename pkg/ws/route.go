package ws

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func Route(context *events.APIGatewayWebsocketProxyRequestContext, body string) {
	var message Message
	err := message.Decode([]byte(body))
	if err != nil {
		fmt.Println("Error decoding message", err)
		return
	}

	fmt.Println("Route ", message.Service)
	switch message.Service {
	case "chat":
		HandleChat(context, message)
		break
	}

}
