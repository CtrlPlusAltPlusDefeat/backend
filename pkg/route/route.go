package route

import (
	"backend/pkg/models"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

// Message Generic wrapper for all websocket messages

func Route(context *events.APIGatewayWebsocketProxyRequestContext, body string) {
	var message models.Wrapper
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
