#!/bin/sh
echo "Starting DynamoDB"
docker-compose up --wait --force-recreate

echo "Creating tables"
aws db create-table --cli-input-json file://db/test-table.json --endpoint-url "http://localhost:8000"

echo "Tables:"
aws db list-tables --endpoint-url "http://localhost:8000"

echo "Compiling code"
./compile-go.sh

sam local start-api --docker-network "backend_local-api-network"
