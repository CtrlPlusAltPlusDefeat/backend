package handlers

import (
	apigateway "backend/pkg/aws-helpers/api-gateway"
	"backend/pkg/models"
	"backend/pkg/ws"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

// DisconnectHandler we want to delete the connection information from db here and return an ok response
func DisconnectHandler(context context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	log.Printf("DisconnectHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)

	con := models.NewContext(context, &req.RequestContext.ConnectionID, &req.RequestContext.DomainName, &req.RequestContext.Stage)
	//todo get all lobbyIds for sessionId
	err := ws.Disconnect(con)

	if err != nil {
		return apigateway.Response{}, err
	}

	return apigateway.OkResponse(), nil
}
