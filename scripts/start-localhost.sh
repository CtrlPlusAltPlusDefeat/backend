#!/bin/sh
echo "Starting DynamoDB"
docker-compose up --wait --force-recreate

echo "Creating tables"
aws dynamodb create-table --cli-input-json file://../db/Connection.json --endpoint-url "http://localhost:8000"

echo "Tables:"
aws dynamodb list-tables --endpoint-url "http://localhost:8000"

#sam local start-api --docker-network "backend_local-api-network"
go run ../pkg/local-main.go 443
