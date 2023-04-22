package ws

import (
	"backend/pkg/aws-helpers"
	"backend/pkg/db"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi/types"
	"github.com/aws/smithy-go"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

var LocalConnections = make(map[string]*websocket.Conn)

func getClient() *apigatewaymanagementapi.Client {
	return apigatewaymanagementapi.NewFromConfig(aws_helpers.GetConfig())
}

// Send sends the provided data to the provided Amazon API Gateway connection ID. A common failure scenario which
// results in an error is if the connection ID is no longer valid. This can occur when a client disconnected from the
// Amazon API Gateway endpoint but the disconnect AWS Lambda was not invoked as it is not guaranteed to be invoked when
// clients disconnect.
func Send(ctx context.Context, id string, data []byte) error {
	// check env vars to see if its running locally if so we can pull connection from map and write
	// else we just use apigateway
	local := os.Getenv("LOCAL_WEBSOCKET_SERVER")
	if local == "1" {
		return writeMessage(id, data)
	}

	//use apigateway when not local
	_, err := getClient().PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
		Data:         data,
		ConnectionId: aws.String(id),
	})
	return handleError(err, id)
}

// writeMessage this is only used when running locally
func writeMessage(connectionId string, data []byte) error {
	connection := LocalConnections[connectionId]
	if connection == nil {
		return deleteConnection(connectionId)
	}
	err := connection.WriteMessage(1, data)
	isClosed := websocket.IsCloseError(err)
	if isClosed {
		return deleteConnection(connectionId)
	}
	return err
}

// handleError is a convenience function for taking action for a given error value. The function handles nil errors as a
// convenience to the caller. If a nil error is provided, the error is immediately returned. The function may return an
// error from the handling action, such as deleting the id from the cache, if that action results in an error.
func handleError(err error, id string) error {
	if err == nil {
		return err
	}

	// Casting to the awserr.Error type will allow you to inspect the error code returned by the service in code. The
	// error code can be used to switch on context specific functionality.

	if err != nil {
		var serializationError *smithy.SerializationError
		if errors.As(err, &serializationError) {
			log.Printf("SerializationError, delete stale connection details from cache: %s", id)
			return deleteConnection(id)
		}
	}

	if err != nil {
		var gone *types.GoneException
		if errors.As(err, &gone) {
			log.Printf("GoneException, delete stale connection details from cache: %s", id)
			return deleteConnection(id)
		}
	}
	return err
}

// handle disconnecting connections
func deleteConnection(id string) error {
	log.Printf("Deleting connection Id: %s", id)
	return db.Connection.GetClient().Remove(db.ConnectionItem{ConnectionId: id})
}
