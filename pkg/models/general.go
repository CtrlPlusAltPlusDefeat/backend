package models

import "github.com/aws/aws-lambda-go/events"

type SocketData struct {
	RequestContext *events.APIGatewayWebsocketProxyRequestContext
	Message        Wrapper
	SessionId      *string
}
