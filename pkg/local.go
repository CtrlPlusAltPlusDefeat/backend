package main

import (
	"backend/pkg/handlers"
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
		// Listen for an incoming connection.
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		handler, err := handlers.ConnectHandler(context.TODO(), &events.APIGatewayWebsocketProxyRequest{
			RequestContext: events.APIGatewayWebsocketProxyRequestContext{ConnectionID: "", RequestID: ""},
		})
		if err != nil {
			return
		}


	}
}
