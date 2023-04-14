package main

import (
	"backend/pkg/handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handlers.DefaultHandler)
}
