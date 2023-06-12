package context

import (
	"backend/pkg/models"
	"backend/pkg/models/game"
	"context"
)

type Context struct {
	value       context.Context
	route       *models.Route
	connection  *ConnectionContext
	gameSession *game.Session
	sessionId   *string
	lobby       *models.Lobby
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

func (c *Context) Value() context.Context {
	return c.value
}

func (c *Context) SessionId() *string {
	return c.sessionId
}

func (c *Context) Lobby() *models.Lobby {
	return c.lobby
}

func (c *Context) Route() *models.Route {
	return c.route
}

func (c *Context) LobbyId() *string {
	return &c.lobby.LobbyId
}

func (c *Context) ConnectionId() *string {
	return c.connection.id
}

func (c *Context) ConnectionHost() *string {
	return c.connection.host
}

func (c *Context) ConnectionPath() *string {
	return c.connection.path
}

func (c *Context) ForConnection(id *string) *Context {
	c.connection.id = id
	return c
}

func (c *Context) ForSession(sessionId *string) *Context {
	c.sessionId = sessionId
	return c
}

func (c *Context) ForLobby(lobby *models.Lobby) *Context {
	c.lobby = lobby
	return c
}

func (c *Context) ForRoute(route *models.Route) *Context {
	c.route = route
	return c
}

func (c *Context) ForGameSession(gameSession *game.Session) *Context {
	c.gameSession = gameSession
	return c
}

func (c *Context) GameSession() *game.Session {
	return c.gameSession
}
