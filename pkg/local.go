package main

import (
	"backend/pkg/handlers"
	"bufio"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	//Listen for incoming connections.
	listener, err := net.Listen("tcp", ":"+arguments[1])
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Listening on " + arguments[1])

	//Close the listener when the application closes.
	defer closeListener(listener)
	handleConnection(listener)
}

func closeListener(listener net.Listener) {
	err := listener.Close()
	if err != nil {
		panic(err)
	}
}

func handleConnection(listener net.Listener) {
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
	connectionId := fmt.Sprintf("id-%s-%s", connection.LocalAddr(), connection.RemoteAddr())
	_, err := handlers.ConnectHandler(context.TODO(), &events.APIGatewayWebsocketProxyRequest{
		RequestContext: events.APIGatewayWebsocketProxyRequestContext{ConnectionID: connectionId, RequestID: ""},
	})
	if err != nil {
		return
	}

	// Make a buffer to hold incoming data.
	scanner := bufio.NewScanner(connection)
	// todo path this to handlers.DefaultHandler
	// todo look into programmatically writing to socket and apigateway connection
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("[server/Connection][%s] Message Received:%s", connectionId, message)

		response := fmt.Sprintf("[server/Connection][%s] Received message. Trying to route it", connectionId)
		_, err := connection.Write([]byte(response))
		if err != nil {
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("[server/Connection] Error reading: ", err.Error())
	}
}
