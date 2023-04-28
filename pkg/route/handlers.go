package route

// References
// https://github.com/aws-samples/apigateway-websockets-golang
// https://github.com/aws-samples/simple-websockets-chat-app/blob/master/onconnect/app.js

import (
	apigateway "backend/pkg/aws-helpers/api-gateway"
	"backend/pkg/ws"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

// ConnectHandler we want to store the connection information in db here and return an ok response
func ConnectHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	log.Printf("ConnectHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)

	err := ws.Connect(req.RequestContext.ConnectionID)
	if err != nil {
		return apigateway.Response{}, err
	}

	return apigateway.OkResponse(), nil
}

// DisconnectHandler we want to delete the connection information from db here and return an ok response
func DisconnectHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	log.Printf("DisconnectHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)

	err := ws.Disconnect(&req.RequestContext.ConnectionID)
	if err != nil {
		return apigateway.Response{}, err
	}

	return apigateway.OkResponse(), nil
}

// DefaultHandler this is where all normal requests will come in
func DefaultHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	log.Printf("DefaultHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)
	log.Printf("msg %s", req.Body)

	Route(&req.RequestContext, req.Body)
	return apigateway.OkResponse(), nil
}
