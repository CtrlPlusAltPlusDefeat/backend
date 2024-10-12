# Running Locally

1. Install docker
2. Make sure to setup an aws profile locally when running dynamo db
3. Install air using the command `make install-air`
4. Then run `make watch` to run with hot reloading

# Player Service

## Sessions

When a player connects to the websocket the server will expect a session request

### Client Request:

The client will send one of two request:
Create a new session

```json
{
  "service": "player",
  "action": "create-session",
  "data": {}
}
```

Use an existing session, this should check no other connections are using it as well as validating it's a `guid`

```json
{
  "service": "player",
  "action": "use-session",
  "data": {
    "sessionId": "\"session\""
  }
}
```

### Server Response:

```json
{
  "service": "player",
  "action": "set-session",
  "data": {
    "sessionId": "\"session\""
  }
}
```

---

# Lobby Service

## Objects

### PlayerObject

```json
{
  "id": "",
  "name": "",
  "isAdmin": "",
  "points": ""
}
```

isAdmin will be used on client side to show the lobby settings, allowing the player to pick game etc. All changes will
be verified server side as well

### LobbyObject

```json
{
  "id": "",
  "players": [
    "<PlayerObject>"
  ],
  "selectedGame": {}
}
```

## Client

### Create

```json
{
  "service": "lobby",
  "action": "create",
  "data": {}
}
```

### Join

```json
{
  "service": "lobby",
  "action": "join",
  "data": {
    "sessionId": ""
  }
}
```

### Get

```json
{
  "service": "lobby",
  "action": "get",
  "data": {
    "sessionId": ""
  }
}
```

### Set Name

```json
{
  "service": "lobby",
  "action": "set-name",
  "data": {
    "text": "",
    "lobbyId": ""
  }
}
```

## Server

### Joined

The response for Lobby create/Lobby join:

```json
{
  "service": "lobby",
  "action": "joined",
  "data": {
    "sessionId": ""
  }
}
```

### Get

```json
{
  "service": "lobby",
  "action": "get",
  "data": {
    "player": "<Player>",
    "lobby": "<Lobby>"
  }
}
```

If they are new then they will have null for the name. The FE should then force the player to enter a name

### Player Joined

```json
{
  "service": "lobby",
  "action": "player-joined",
  "data": {
    "player": "<PlayerObject>"
  }
}
```

### Player Left

```json
{
  "service": "lobby",
  "action": "player-left",
  "data": {
    "playerId": "<PlayerObject>"
  }
}
```

### Player Name Change

```json
{
  "service": "lobby",
  "action": "name-change",
  "data": {
    "player": "<PlayerObject>"
  }
}
```

# Chat Service

Sending chat messages using the websocket requires data in the following format

### Client Request:

```json
{
  "service": "chat",
  "action": "send",
  "data": {
    "text": "\"Hello\""
  }
}
```

### Server Response:

todo - need to change connectionId to sessionId

```json
{
  "service": "chat",
  "action": "received",
  "data": {
    "text": "\"Hello\"",
    "connectionId": "475955b2-8aba-49c9-b957-58f719aefbc9"
  }
}
```

---