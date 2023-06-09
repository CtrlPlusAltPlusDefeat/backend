package game

import (
	"backend/pkg/models"
)

type EventType string

type ControllerHandler func(ctx *models.Context, controller *Controller)

type Middleware func(handler ControllerHandler) ControllerHandler

const (
	OnTurnEnd     EventType = "onturnend"
	OnPlayerInput EventType = "onplayerinput"
)

type Controller struct {
	handlers  map[EventType]ControllerHandler
	GameState *State
}

func NewController(state *State) *Controller {
	return &Controller{
		handlers:  make(map[EventType]ControllerHandler),
		GameState: state,
	}
}

func (c *Controller) AddHandler(eventType EventType, handler ControllerHandler, middleware ...Middleware) {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	c.handlers[eventType] = handler
}

func (c *Controller) ExecuteHandlers(eventType EventType, ctx *models.Context) {
	handler := c.handlers[eventType]
	if handler != nil {
		handler(ctx, c)
	}
}
