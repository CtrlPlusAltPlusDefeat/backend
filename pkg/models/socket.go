package models

import "github.com/aws/aws-lambda-go/events"

type SocketData struct {
	RequestContext *events.APIGatewayWebsocketProxyRequestContext
	Message        Wrapper
	SessionId      *string
}

type Connection struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
	SessionId    string `dynamodbav:"SessionId"`
}
