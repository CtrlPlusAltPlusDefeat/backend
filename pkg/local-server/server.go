package local_server

import (
	"backend/pkg/handlers"
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"net"
	"os"
)

var connections = make(map[string]net.Conn)

func HandleConnection(listener net.Listener) {
	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		acceptConnection(connection)
	}
}

func acceptConnection(connection net.Conn) {
	//new connection
	connectionId := uuid.New().String()
	connections[connectionId] = connection

	_, err := handlers.ConnectHandler(context.TODO(), &events.APIGatewayWebsocketProxyRequest{
		RequestContext: events.APIGatewayWebsocketProxyRequestContext{ConnectionID: connectionId, RequestID: ""},
	})
	if err != nil {
		return
	}

	// Make a buffer to hold incoming data.
	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		message := scanner.Text()
		_, err = handlers.DefaultHandler(context.TODO(), &events.APIGatewayWebsocketProxyRequest{
			RequestContext: events.APIGatewayWebsocketProxyRequestContext{ConnectionID: connectionId},
			Body:           message,
		})
		if err != nil {
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[server/Connection] Error reading: ", err.Error())
	}
}

func WriteMessage(connectionId string, data []byte) {
	connection := connections[connectionId]
	_, err := connection.Write(data)
	if err != nil {
		fmt.Println("Error writing message", err)
		return
	}
}
