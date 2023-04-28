package models

import (
	"context"
)

type Context struct {
	value context.Context

	connection *ConnectionContext

	sessionId *string
	lobbyId   *string
}

type ConnectionContext struct {
	id   *string
	host *string
	path *string
}

func NewContext(value context.Context, connectionId *string, connectionHost *string, connectionPath *string) *Context {
	return &Context{
		value: value,
		connection: &ConnectionContext{
			id:   connectionId,
			host: connectionHost,
			path: connectionPath,
		},
	}
}

func (c Context) Value() context.Context {
	return c.value
}

func (c Context) SessionId() *string {
	return c.sessionId
}

func (c Context) LobbyId() *string {
	return c.lobbyId
}

func (c Context) ConnectionId() *string {
	return c.connection.id
}

func (c Context) ConnectionHost() *string {
	return c.connection.host
}

func (c Context) ConnectionPath() *string {
	return c.connection.path
}

func (c Context) ForConnection(id *string) *Context {
	return &Context{
		value:     c.value,
		lobbyId:   c.lobbyId,
		sessionId: c.sessionId,
		connection: &ConnectionContext{
			id:   id,
			host: c.connection.host,
			path: c.connection.path,
		},
	}
}

func (c Context) ForSession(id *string) *Context {
	return &Context{
		value:     c.value,
		lobbyId:   c.lobbyId,
		sessionId: id,
		connection: &ConnectionContext{
			id:   c.connection.id,
			host: c.connection.host,
			path: c.connection.path,
		},
	}
}

func (c Context) ForLobby(id *string) *Context {
	return &Context{
		value:     c.value,
		lobbyId:   id,
		sessionId: c.sessionId,
		connection: &ConnectionContext{
			id:   c.connection.id,
			host: c.connection.host,
			path: c.connection.path,
		},
	}
}
