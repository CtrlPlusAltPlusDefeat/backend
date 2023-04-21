package route

// References
// https://github.com/aws-samples/apigateway-websockets-golang
// https://github.com/aws-samples/simple-websockets-chat-app/blob/master/onconnect/app.js

import (
	apigateway "backend/pkg/aws-helpers/api-gateway"
	"backend/pkg/db"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

// ConnectHandler we want to store the connection information in db here and return an ok response
func ConnectHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	fmt.Printf("ConnectHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)

	err := db.GetConnectionDb().Add(db.Connection{ConnectionId: req.RequestContext.ConnectionID})
	if err != nil {
		return apigateway.Response{}, err
	}

	return apigateway.OkResponse(), nil
}

// DisconnectHandler we want to delete the connection information from db here and return an ok response
func DisconnectHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	fmt.Printf("DisconnectHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)

	err := db.GetConnectionDb().Remove(db.Connection{ConnectionId: req.RequestContext.ConnectionID})
	if err != nil {
		return apigateway.Response{}, err
	}

	return apigateway.OkResponse(), nil
}

// DefaultHandler this is where all normal requests will come in
func DefaultHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	fmt.Printf("DefaultHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)

	Route(&req.RequestContext, req.Body)
	return apigateway.OkResponse(), nil
}
