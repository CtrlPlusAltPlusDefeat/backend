#!/bin/sh
if [ -z "$PROJECT_ROOT" ]; then
  export PROJECT_ROOT="."
fi

echo "Executing from $PROJECT_ROOT"

GOOS=linux GOARCH=amd64 go build -o "$PROJECT_ROOT/dist/disconnect" "$PROJECT_ROOT/pkg/handlers/disconnect/disconnect.go"
GOOS=linux GOARCH=amd64 go build -o "$PROJECT_ROOT/dist/connect" "$PROJECT_ROOT/pkg/handlers/connect/connect.go"
GOOS=linux GOARCH=amd64 go build -o "$PROJECT_ROOT/dist/default" "$PROJECT_ROOT/pkg/handlers/default/default.go"

echo "compiled code"