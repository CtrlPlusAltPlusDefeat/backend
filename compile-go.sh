#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o dist/test-handler src/test-handler.go
GOOS=linux GOARCH=amd64 go build -o dist/disconnect src/disconnect/disconnect.go
GOOS=linux GOARCH=amd64 go build -o dist/connect src/connect/connect.go
echo "compiled code"
