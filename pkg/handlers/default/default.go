package main

import (
	"backend/pkg/db"
	"backend/pkg/route"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
*
init is called when lambda starts up,
spin up a dynamodb client and inject into db package
*/
func init() {
	db.Configure()
}

func main() {
	lambda.Start(route.DefaultHandler)
}
