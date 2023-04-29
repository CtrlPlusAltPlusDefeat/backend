package handlers

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
