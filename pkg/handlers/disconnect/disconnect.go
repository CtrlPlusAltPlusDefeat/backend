package main

import (
	"backend/pkg/route"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(route.DisconnectHandler)
}
