package models

import (
	"context"
)

type Context struct {
	value      context.Context
	route      *Route
	connection *ConnectionContext

	sessionId *string
	lobby     *Lobby
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

func (c Context) Lobby() *Lobby {
	return c.lobby
}

func (c Context) Route() *Route {
	return c.route
}

func (c Context) LobbyId() *string {
	return &c.lobby.LobbyId
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
		lobby:     c.lobby,
		route:     c.route,
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
		lobby:     c.lobby,
		sessionId: id,
		route:     c.route,
		connection: &ConnectionContext{
			id:   c.connection.id,
			host: c.connection.host,
			path: c.connection.path,
		},
	}
}

func (c Context) ForLobby(lobby *Lobby) *Context {
	return &Context{
		value:     c.value,
		lobby:     lobby,
		route:     c.route,
		sessionId: c.sessionId,
		connection: &ConnectionContext{
			id:   c.connection.id,
			host: c.connection.host,
			path: c.connection.path,
		},
	}
}

func (c Context) ForRoute(route *Route) *Context {
	return &Context{
		value:     c.value,
		lobby:     c.lobby,
		route:     route,
		sessionId: c.sessionId,
		connection: &ConnectionContext{
			id:   c.connection.id,
			host: c.connection.host,
			path: c.connection.path,
		},
	}
}
