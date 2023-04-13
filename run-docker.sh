#!/bin/sh
rm -rf ./db/shared-local-instance.db
docker-compose rm -f
docker-compose pull
docker-compose up --build -d
aws dynamodb create-table --cli-input-json file://db/test-table.json --endpoint-url http://localhost:8000
sam build --use-container
sam local start-api --port 3001 --skip-pull-image