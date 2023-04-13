#!/bin/sh
export NETWORK_NAME="local-api-network"
export CONTAINER_NAME="dynamo-local"
export CONTAINER_ENDPOINT="http://localhost:8000"

echo "Starting DynamoDB"
docker network create "$NETWORK_NAME"
docker run -d -p 8000:8000 --network="$NETWORK_NAME" --name "$CONTAINER_NAME" amazon/dynamodb-local

echo "Creating tables"
aws dynamodb create-table --cli-input-json file://db/test-table.json --endpoint-url "$CONTAINER_ENDPOINT"

echo "Tables:"
aws dynamodb list-tables --endpoint-url "$CONTAINER_ENDPOINT"

echo "Compiling code"
./compile-go.sh

sam local start-api --docker-network "$NETWORK_NAME"

echo "cleaning up"
docker ps -aq | xargs docker stop | xargs docker rm
docker network rm "$NETWORK_NAME"
