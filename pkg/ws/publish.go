package ws

import (
	"backend/pkg/aws-helpers"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

func getClient() *apigatewaymanagementapi.Client {
	return apigatewaymanagementapi.NewFromConfig(aws_helpers.GetConfig())
}

// Send sends the provided data to the provided Amazon API Gateway connection ID. A common failure scenario which
// results in an error is if the connection ID is no longer valid. This can occur when a client disconnected from the
// Amazon API Gateway endpoint but the disconnect AWS Lambda was not invoked as it is not guaranteed to be invoked when
// clients disconnect.
func Send(ctx context.Context, id string, data []byte) error {
	_, err := getClient().PostToConnection(ctx, &apigatewaymanagementapi.PostToConnectionInput{
		Data:         data,
		ConnectionId: aws.String(id),
	})
	return err
}
