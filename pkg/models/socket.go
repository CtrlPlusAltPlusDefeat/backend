package models

import (
	"context"
)

type Connection struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
	SessionId    string `dynamodbav:"SessionId"`
}

type Data struct {
	Message Wrapper
}

type Context struct {
	Value context.Context

	Connection *ConnectionContext

	SessionId *string
	LobbyId   *string
}

type ConnectionContext struct {
	Id   *string
	Host *string
	Path *string
}

func (c Context) ForConnection(id *string) Context {
	return Context{
		Value:     c.Value,
		LobbyId:   c.LobbyId,
		SessionId: c.SessionId,
		Connection: &ConnectionContext{
			Id:   id,
			Host: c.Connection.Host,
			Path: c.Connection.Path,
		},
	}
}
