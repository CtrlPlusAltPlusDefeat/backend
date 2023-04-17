package ws

import (
	"backend/pkg/aws-helpers"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/gorilla/websocket"
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
		writeMessage(id, data)
		return nil
	}
	_, err := getClient().PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
		Data:         data,
		ConnectionId: aws.String(id),
	})
	return err
}

func writeMessage(connectionId string, data []byte) {
	connection := LocalConnections[connectionId]
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