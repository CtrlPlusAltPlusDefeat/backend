#!/bin/sh
GOOS=linux GOARCH=amd64 go build -o dist/test-handler test-handler.go
echo "compiled code"