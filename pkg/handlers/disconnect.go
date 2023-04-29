package handlers

import (
	apigateway "backend/pkg/aws-helpers/api-gateway"
	"backend/pkg/ws"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

// DisconnectHandler we want to delete the connection information from db here and return an ok response
func DisconnectHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	log.Printf("DisconnectHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)

	err := ws.Disconnect(&req.RequestContext.ConnectionID)

	if err != nil {
		return apigateway.Response{}, err
	}

	return apigateway.OkResponse(), nil
}
