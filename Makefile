install-air:
	go install github.com/cosmtrek/air@latest

air:
	~/.air -c .air.toml

build:
	go build -o ./tmp/main ./pkg/local-main.go

dynamodb:
	rm -rf ./docker
	echo "Starting DynamoDB"
	docker-compose up --wait --force-recreate
	echo "Creating tables"
	AWS_PAGER="" aws dynamodb create-table --cli-input-json file://./db/Connection.json --endpoint-url "http://localhost:8000"
	AWS_PAGER="" aws dynamodb create-table --cli-input-json file://./db/Lobby.json --endpoint-url "http://localhost:8000"
	AWS_PAGER="" aws dynamodb create-table --cli-input-json file://./db/LobbyPlayer.json --endpoint-url "http://localhost:8000"
	AWS_PAGER="" aws dynamodb create-table --cli-input-json file://./db/LobbyChat.json --endpoint-url "http://localhost:8000"
	AWS_PAGER="" aws dynamodb create-table --cli-input-json file://./db/GameSession.json --endpoint-url "http://localhost:8000"

	echo "Tables:"
	aws dynamodb list-tables --endpoint-url "http://localhost:8000"

watch:
	make dynamodb && make air

connect-locally:
	wscat -c  ws://localhost:8080

lint:
	golangci-lint run ./pkg

add-data:
	aws dynamodb put-item \
	  --table-name Lobby \
	  --item file://./db/Data/Lobby.json \
	  --return-consumed-capacity TOTAL \
	  --return-item-collection-metrics SIZE \
	  --endpoint-url "http://localhost:8000"

	aws dynamodb put-item \
	  --table-name LobbyPlayer \
	  --item file://./db/Data/LobbyPlayer1.json \
	  --return-consumed-capacity TOTAL \
	  --return-item-collection-metrics SIZE \
	  --endpoint-url "http://localhost:8000"

	aws dynamodb put-item \
	  --table-name LobbyPlayer \
	  --item file://./db/Data/LobbyPlayer2.json \
	  --return-consumed-capacity TOTAL \
	  --return-item-collection-metrics SIZE \
	  --endpoint-url "http://localhost:8000"

	aws dynamodb put-item \
	  --table-name LobbyPlayer \
	  --item file://./db/Data/LobbyPlayer3.json \
	  --return-consumed-capacity TOTAL \
	  --return-item-collection-metrics SIZE \
	  --endpoint-url "http://localhost:8000"