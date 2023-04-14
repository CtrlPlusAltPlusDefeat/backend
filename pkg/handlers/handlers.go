package handlers

// References
// https://github.com/aws-samples/apigateway-websockets-golang
// https://github.com/aws-samples/simple-websockets-chat-app/blob/master/onconnect/app.js

import (
	"backend/pkg/aws-helpers/api-gateway"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

// ConnectHandler we want to store the connection information in dynamodb here and return an ok response
func ConnectHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (api_gateway.Response, error) {
	fmt.Println("requestId", req.RequestContext.RequestID)
	fmt.Println("connectionId", req.RequestContext.ConnectionID)

	return api_gateway.OkResponse(), nil
}

// DisconnectHandler we want to delete the connection information from dynamodb here and return an ok response
func DisconnectHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (api_gateway.Response, error) {
	fmt.Println("requestId", req.RequestContext.RequestID)
	fmt.Println("connectionId", req.RequestContext.ConnectionID)

	return api_gateway.OkResponse(), nil
}

// DefaultHandler this is where all responses will come into
func DefaultHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (api_gateway.Response, error) {
	fmt.Println("requestId", req.RequestContext.RequestID)
	fmt.Println("connectionId", req.RequestContext.ConnectionID)

	return api_gateway.OkResponse(), nil
}
