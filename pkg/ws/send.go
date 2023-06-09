package ws

import (
	awshelpers "backend/pkg/aws-helpers"
	"backend/pkg/db"
	"backend/pkg/models"
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi/types"
	"github.com/aws/smithy-go"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
)

var LocalConnections = make(map[string]*websocket.Conn)

func getClient(context *models.Context) *apigatewaymanagementapi.Client {
	callbackURL := url.URL{
		Scheme: "https",
		Host:   *context.ConnectionHost(),
		Path:   *context.ConnectionPath(),
	}

	log.Println("Creating API Gateway client for callback URL: ", callbackURL.String())

	return apigatewaymanagementapi.NewFromConfig(awshelpers.GetConfig(), func(o *apigatewaymanagementapi.Options) {
		o.EndpointResolver = apigatewaymanagementapi.EndpointResolverFromURL(callbackURL.String())
	})
}

func SendToLobby(context *models.Context, route *models.Route, message interface{}) error {
	players, err := db.LobbyPlayer.GetPlayers(&context.Lobby().LobbyId)

	if err != nil {
		return err
	}

	log.Println("Sending to ", len(players), " players")

	for index, p := range players {
		log.Println("Sending to player ", index)

		err = Send(context.ForConnection(&p.ConnectionId), route, message)

		if err != nil {
			log.Printf("sendToLobby error sending to %s ", p.ConnectionId)
		}
	}

	return nil
}

// Send sends the provided data to the provided Amazon API Gateway connection ID. A common failure scenario which
// results in an error is if the connection ID is no longer valid. This can occur when a client disconnected from the
// Amazon API Gateway endpoint but the disconnect AWS Lambda was not invoked as it is not guaranteed to be invoked when
// clients disconnect.
func Send(context *models.Context, route *models.Route, message any) error {

	value, _ := json.Marshal(message)
	wrapper := models.Wrapper{Service: *route.Service(), Action: *route.Action(), Data: string(value)}
	data, _ := json.Marshal(wrapper)

	log.Printf("Sending: %s", data)

	// check env vars to see if its running locally if so we can pull connection from map and write
	// else we just use apigateway
	local := os.Getenv("LOCAL_WEBSOCKET_SERVER")
	if local == "1" {
		return writeMessage(context.ConnectionId(), data)
	}

	//use apigateway when not local
	_, err := getClient(context).PostToConnection(context.Value(), &apigatewaymanagementapi.PostToConnectionInput{
		Data:         data,
		ConnectionId: context.ConnectionId(),
	})
	return handleError(err, context.ConnectionId())
}

// writeMessage this is only used when running locally
func writeMessage(connectionId *string, data []byte) error {
	connection := LocalConnections[*connectionId]
	if connection == nil {
		return db.Connection.Remove(connectionId)
	}
	err := connection.WriteMessage(1, data)
	isClosed := websocket.IsCloseError(err)
	if isClosed {
		return db.Connection.Remove(connectionId)
	}
	return err
}

// handleError is a convenience function for taking action for a given error value. The function handles nil errors as a
// convenience to the caller. If a nil error is provided, the error is immediately returned. The function may return an
// error from the handling action, such as deleting the id from the cache, if that action results in an error.
func handleError(err error, id *string) error {
	if err == nil {
		return err
	}

	if err != nil {
		var serializationError *smithy.SerializationError
		if errors.As(err, &serializationError) {
			log.Printf("SerializationError, delete stale connection details from cache: %s", id)
			return db.Connection.Remove(id)
		}
	}

	if err != nil {
		var gone *types.GoneException
		if errors.As(err, &gone) {
			log.Printf("GoneException, delete stale connection details from cache: %s", id)
			return db.Connection.Remove(id)
		}
	}
	return err
}
