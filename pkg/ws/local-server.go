package ws

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
)

var connections = make(map[string]*websocket.Conn)

func HandleConnection(conn *websocket.Conn) {
	//new connection
	connectionId := uuid.New().String()
	connections[connectionId] = conn

	_, err := ConnectHandler(context.TODO(), &events.APIGatewayWebsocketProxyRequest{
		RequestContext: events.APIGatewayWebsocketProxyRequestContext{ConnectionID: connectionId, RequestID: ""},
	})
	if err != nil {
		return
	}

	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		log.Println(string(p))

		_, err = DefaultHandler(context.TODO(), &events.APIGatewayWebsocketProxyRequest{
			RequestContext: events.APIGatewayWebsocketProxyRequestContext{ConnectionID: connectionId},
			Body:           string(p),
		})
	}
}

func WriteMessage(connectionId string, data []byte) {
	connection := connections[connectionId]
	if connection == nil {
		_, err := DisconnectHandler(context.TODO(), &events.APIGatewayWebsocketProxyRequest{
			RequestContext: events.APIGatewayWebsocketProxyRequestContext{ConnectionID: connectionId},
		})
		if err != nil {
			fmt.Println("Error writing message", err)
		}
		return
	}
	err := connection.WriteMessage(1, data)
	if err != nil {
		fmt.Println("Error writing message", err)
		return
	}
}
