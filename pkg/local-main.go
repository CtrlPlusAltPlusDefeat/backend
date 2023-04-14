package main

import (
	localserver "backend/pkg/ws"
	"fmt"
	"net"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	// set environment variables
	_ = os.Setenv("LOCAL_WEBSOCKET_SERVER", "1")
	_ = os.Setenv("DYNAMO_DB_URL", "http://localhost:8000")

	//Listen for incoming connections.
	listener, err := net.Listen("tcp", ":"+arguments[1])
	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Listening on " + arguments[1])

	//Close the listener when the application closes.
	defer closeListener(listener)
	localserver.HandleConnection(listener)
}

func closeListener(listener net.Listener) {
	err := listener.Close()
	if err != nil {
		panic(err)
	}
}
