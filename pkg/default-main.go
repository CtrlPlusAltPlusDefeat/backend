package main

import (
	"backend/pkg/db"
	"backend/pkg/handlers"
	"backend/pkg/routes"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
*
init is called when lambda starts up,
spin up a dynamodb client and inject into db package
*/
func init() {
	db.Configure()
	routes.Configure()
}

func main() {
	lambda.Start(handlers.DefaultHandler)
}
