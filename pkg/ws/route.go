package ws

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

// Message Generic wrapper for all websocket messages
type Message struct {
	Service string `json:"service"`
	Action  string `json:"action"`
	Data    string `json:"data"`
}

func (message *Message) Encode() ([]byte, error) {
	return json.Marshal(message)
}

func (message *Message) Decode(data []byte) error {
	return json.Unmarshal(data, message)
}

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
