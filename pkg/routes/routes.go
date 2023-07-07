package routes

import (
	"backend/pkg/models"
	"backend/pkg/models/context"
	"backend/pkg/services"
	"log"
)

type (
	Middleware func(handler Handler) Handler

	Handler func(context *context.Context, data *models.Data) error
)

var (
	handlers = make(map[string]Handler)
)

func Configure() {
	add("player|create-session", services.CreateSession, ErrorCommunicateMiddleware)
	add("player|use-session", services.UseSession, ErrorCommunicateMiddleware)

	add("lobby|create", services.CreateLobby, ErrorCommunicateMiddleware, SessionMiddleware)
	add("lobby|join", services.JoinLobby, ErrorCommunicateMiddleware, SessionMiddleware, LobbyMiddleware)
	add("lobby|leave", services.LeaveLobby, ErrorCommunicateMiddleware, SessionMiddleware, LobbyMiddleware)
	add("lobby|load-game", services.LoadGame, ErrorCommunicateMiddleware, SessionMiddleware, LobbyMiddleware)

	add("chat|send", services.SendChat, ErrorCommunicateMiddleware, SessionMiddleware, LobbyMiddleware)
	add("chat|load", services.LoadChat, ErrorCommunicateMiddleware, SessionMiddleware, LobbyMiddleware)

	add("game|get-state", services.GetState, ErrorCommunicateMiddleware, SessionMiddleware, LobbyMiddleware, GameSessionMiddleware)
	add("game|player-action", services.PlayerAction, ErrorCommunicateMiddleware, SessionMiddleware, LobbyMiddleware, GameSessionMiddleware)
	add("game|swap-teams", services.SwapTeam, ErrorCommunicateMiddleware, SessionMiddleware, LobbyMiddleware, GameSessionMiddleware)
}

func add(route string, handler Handler, middleware ...Middleware) {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	handlers[route] = handler
}

func Execute(context *context.Context, data *models.Data) {
	route := data.Route().Value()
	log.Printf("Beginning Invoking '%s'", route)
	context = context.ForRoute(data.Route())
	if handler, exists := handlers[route]; exists {
		err := handler(context, data)
		if err != nil {
			log.Printf("Error Invoking '%s': %s", route, err)
		}
	} else {
		log.Printf("Error Invoking '%s': Couldn't find route.", route)
	}
}
