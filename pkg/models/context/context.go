package context

import (
	"backend/pkg/game"
	"backend/pkg/models"
	"context"
)

type BeforeSend func(*Context, *models.Player, any) (any, error)

type Context struct {
	value       context.Context
	route       *models.Route
	connection  *ConnectionContext
	sessionId   *string
	gameSession *game.Session
	lobby       *models.Lobby
	beforeSend  *BeforeSend
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

func (c *Context) GameId() int {
	settings, err := c.lobby.Settings.Decode()
	if err != nil {
		return -1
	}
	return int(settings.GameId)
}

func (c *Context) Route() *models.Route {
	return c.route
}

func (c *Context) LobbyId() *string {
	return &c.lobby.LobbyId
}

func (c *Context) GameSession() *game.Session {
	return c.gameSession
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

func (c *Context) BeforeSend() *BeforeSend {
	return c.beforeSend
}

func (c *Context) ForConnection(id *string) *Context {
	n := c.duplicate()
	n.connection.id = id

	return n
}

func (c *Context) ForSession(sessionId *string) *Context {
	n := c.duplicate()
	n.sessionId = sessionId

	return n
}

func (c *Context) ForLobby(lobby *models.Lobby) *Context {
	n := c.duplicate()
	n.lobby = lobby

	return n
}

func (c *Context) ForRoute(route *models.Route) *Context {
	n := c.duplicate()
	n.route = route

	return n
}

func (c *Context) ForGameSession(gameSession *game.Session) *Context {
	n := c.duplicate()
	n.gameSession = gameSession

	return n
}

func (c *Context) ForBeforeSend(handler *BeforeSend) *Context {
	n := c.duplicate()
	n.beforeSend = handler
	return n
}

func (c *Context) duplicate() *Context {
	return &Context{
		value: c.value,
		connection: &ConnectionContext{
			id:   c.connection.id,
			host: c.connection.host,
			path: c.connection.path,
		},
		route:       c.route,
		gameSession: c.gameSession,
		sessionId:   c.sessionId,
		lobby:       c.lobby,
		beforeSend:  c.beforeSend,
	}
}
