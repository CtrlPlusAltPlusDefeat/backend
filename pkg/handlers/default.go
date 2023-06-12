package handlers

import (
	apigateway "backend/pkg/aws-helpers/api-gateway"
	"backend/pkg/models"
	customCtx "backend/pkg/models/context"
	"backend/pkg/routes"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

// DefaultHandler this is where all normal requests will come in
func DefaultHandler(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	log.Printf("DefaultHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)
	log.Printf("msg %s", req.Body)

	data, _ := models.NewData(req.Body)
	con := customCtx.NewContext(ctx, &req.RequestContext.ConnectionID, &req.RequestContext.DomainName, &req.RequestContext.Stage)

	routes.Execute(con, data)

	return apigateway.OkResponse(), nil
}
