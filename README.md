# Running Locally

1. Install docker
2. Then run the script start-localhost

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