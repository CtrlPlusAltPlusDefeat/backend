package ws

import (
	"backend/pkg/aws-helpers"
	local_server "backend/pkg/local-server"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"os"
)

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
		fmt.Println("Using Local Websocket server")
		local_server.WriteMessage(id, data)
		return nil
	}
	_, err := getClient().PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
		Data:         data,
		ConnectionId: aws.String(id),
	})
	return err
}
