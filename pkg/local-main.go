package main

import (
	awshelpers "backend/pkg/aws-helpers"
	"backend/pkg/db"
	"backend/pkg/route"
	"backend/pkg/ws"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func init() {
	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	db.DynamoDb = dbClient
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	// set environment variables
	_ = os.Setenv("LOCAL_WEBSOCKET_SERVER", "1")
	_ = os.Setenv("DYNAMO_DB_URL", "http://localhost:8000")

	println("Listening on port http://localhost:8080")
	//Listen for incoming connections.
	http.HandleFunc("/", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	if err != nil {
		log.Println(err)
	}

	handleConnection(ws)
}

func handleConnection(conn *websocket.Conn) {
	//new connection
	connectionId := uuid.New().String()
	ws.LocalConnections[connectionId] = conn

	_, err := route.ConnectHandler(context.TODO(), &events.APIGatewayWebsocketProxyRequest{
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

		_, err = route.DefaultHandler(context.TODO(), &events.APIGatewayWebsocketProxyRequest{
			RequestContext: events.APIGatewayWebsocketProxyRequestContext{ConnectionID: connectionId},
			Body:           string(p),
		})
	}
}
