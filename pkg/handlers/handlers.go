package handlers

// References
// https://github.com/aws-samples/apigateway-websockets-golang
// https://github.com/aws-samples/simple-websockets-chat-app/blob/master/onconnect/app.js

import (
	awshelpers "backend/pkg/aws-helpers"
	apigateway "backend/pkg/aws-helpers/api-gateway"
	"backend/pkg/aws-helpers/db"
	"backend/pkg/service"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// ConnectHandler we want to store the connection information in db here and return an ok response
func ConnectHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	fmt.Println("ConnectHandler")
	fmt.Println("requestId", req.RequestContext.RequestID)
	fmt.Println("connectionId", req.RequestContext.ConnectionID)

	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	connectionTable := db.ConnectionTable{DynamoDbClient: dbClient}

	err := connectionTable.Add(db.Connection{ConnectionId: req.RequestContext.ConnectionID})
	if err != nil {
		return apigateway.Response{}, err
	}

	return apigateway.OkResponse(), nil
}

// DisconnectHandler we want to delete the connection information from db here and return an ok response
func DisconnectHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	fmt.Println("requestId", req.RequestContext.RequestID)
	fmt.Println("connectionId", req.RequestContext.ConnectionID)

	dbClient := dynamodb.NewFromConfig(awshelpers.GetConfig())
	connectionTable := db.ConnectionTable{DynamoDbClient: dbClient}

	err := connectionTable.Remove(req.RequestContext.ConnectionID)
	if err != nil {
		return apigateway.Response{}, err
	}

	return apigateway.OkResponse(), nil
}

// DefaultHandler this is where all normal requests will come in
func DefaultHandler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) (apigateway.Response, error) {
	fmt.Println("DefaultHandler")
	fmt.Println("requestId", req.RequestContext.RequestID)
	fmt.Println("connectionId", req.RequestContext.ConnectionID)

	service.Route(&req.RequestContext, req.Body)

	return apigateway.OkResponse(), nil
}
