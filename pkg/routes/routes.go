package routes

import (
	"backend/pkg/models"
	"backend/pkg/routes/middleware"
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
	add("player|create-session", services.CreateSession, middleware.ErrorCommunicateMiddleware)
	add("player|use-session", services.UseSession, middleware.ErrorCommunicateMiddleware)
	add("chat|send", services.SendChat, middleware.SessionMiddleware)
	add("lobby|create", services.CreateLobby, middleware.SessionMiddleware)
	add("lobby|join", services.JoinLobby, middleware.SessionMiddleware)
	add("lobby|set-name", services.SetLobbyName, middleware.SessionMiddleware)
	add("lobby|get", services.GetLobby, middleware.SessionMiddleware)
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
