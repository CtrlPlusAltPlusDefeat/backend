package routes

import (
	"backend/pkg/models"
	"backend/pkg/services"
	"log"
)

type (
	Middleware func(handler Handler) Handler

	Handler func(context *models.Context, data *models.Data) error
)

var (
	handlers map[string]Handler = make(map[string]Handler)
)

func Configure() {
	add("player|create-session", services.CreateSession, ErrorCommunicateMiddleware)
	add("player|use-session", services.UseSession, ErrorCommunicateMiddleware)
	add("chat|send", services.SendChat, SessionMiddleware)
	add("lobby|create", services.CreateLobby, SessionMiddleware)
	add("lobby|join", services.JoinLobby, SessionMiddleware)
	add("lobby|set-name", services.SetLobbyName, SessionMiddleware)
	add("lobby|get", services.GetLobby, SessionMiddleware)
}

func add(route string, handler Handler, middleware ...Middleware) {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	handlers[route] = handler
}

func Execute(context *models.Context, data *models.Data) {
	route := data.Route().Value()

	log.Printf("Beginning Invoking '%s'", route)

	if handler, exists := handlers[route]; exists {
		err := handler(context, data)

		if err != nil {
			log.Printf("Error Invoking '%s': %s", route, err)
		}
	}
}
