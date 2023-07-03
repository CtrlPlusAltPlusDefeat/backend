package handlers

import (
	"backend/pkg/helpers"
	apigateway "backend/pkg/helpers/api-gateway"
	"backend/pkg/models"
	customCtx "backend/pkg/models/context"
	"backend/pkg/routes"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"time"
)

// DefaultHandler this is where all normal requests will come in
func DefaultHandler(ctx context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	log.Printf("DefaultHandler requestId: %s, connectionId:%s \n\r", req.RequestContext.RequestID, req.RequestContext.ConnectionID)
	log.Printf("msg %s", req.Body)

	data, _ := models.NewData(req.Body)
	defer helpers.TimeTrack(time.Now(), data.Route().Value())
	con := customCtx.NewContext(ctx, &req.RequestContext.ConnectionID, &req.RequestContext.DomainName, &req.RequestContext.Stage)

	start := time.Now()
	routes.Execute(con, data)
	elapsed := time.Since(start)
	log.Printf("Executed %s in %s", data.Route().Value(), elapsed)

	return apigateway.OkResponse(), nil
}
