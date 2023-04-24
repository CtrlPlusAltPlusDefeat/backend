aws dynamodb put-item \
  --table-name Lobby \
  --item file://../db/Data/Lobby.json \
  --return-consumed-capacity TOTAL \
  --return-item-collection-metrics SIZE \
  --endpoint-url "http://localhost:8000"

aws dynamodb put-item \
  --table-name LobbyPlayer \
  --item file://../db/Data/LobbyPlayer1.json \
  --return-consumed-capacity TOTAL \
  --return-item-collection-metrics SIZE \
  --endpoint-url "http://localhost:8000"

aws dynamodb put-item \
  --table-name LobbyPlayer \
  --item file://../db/Data/LobbyPlayer2.json \
  --return-consumed-capacity TOTAL \
  --return-item-collection-metrics SIZE \
  --endpoint-url "http://localhost:8000"

aws dynamodb put-item \
  --table-name LobbyPlayer \
  --item file://../db/Data/LobbyPlayer3.json \
  --return-consumed-capacity TOTAL \
  --return-item-collection-metrics SIZE \
  --endpoint-url "http://localhost:8000"
