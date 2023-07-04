package game

import (
	"backend/pkg/models"
	"context"
)

func NewContext(value context.Context, session *Session, data *models.Data, player *models.Player) *Context {
	return &Context{
		value:   value,
		session: session,
		data:    data,
		player:  player,
	}
}

func (c *Context) Value() context.Context {
	return c.value
}

func (c *Context) Data() *models.Data {
	return c.data
}

func (c *Context) Player() *models.Player {
	return c.player
}

func (c *Context) Session() *Session {
	return c.session
}
