package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

// https://stackoverflow.com/questions/62772813/how-i-can-invoke-locally-my-lambda-using-api-gatewat-connect-and-disconnect-ro
// here we want to remove the connection id to dynamo db
func handler(_ context.Context, req *events.APIGatewayWebsocketProxyRequest) {

}
