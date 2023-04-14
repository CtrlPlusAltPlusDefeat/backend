package main

import (
	"backend/pkg/ws"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(ws.DisconnectHandler)
}
