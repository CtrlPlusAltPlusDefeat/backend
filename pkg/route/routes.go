package route

import (
	"backend/pkg/db"
	"backend/pkg/models"
	"backend/pkg/ws"
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
	add("player|create-session", CreateSession, communicateMiddleware)
	add("player|use-session", UseSession, communicateMiddleware)
	add("chat|send", SendChat, sessionMiddleware)
	add("lobby|create", CreateLobby, sessionMiddleware)
	add("lobby|join", JoinLobby, sessionMiddleware)
	add("lobby|set-name", SetLobbyName, sessionMiddleware)
	add("lobby|get", GetLobby, sessionMiddleware)
}

func add(route string, handler Handler, middleware ...Middleware) {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	handlers[route] = handler
}

func sessionMiddleware(next Handler) Handler {
	return func(context *models.Context, data *models.Data) error {
		res, err := db.Connection.Get(context.ConnectionId())

		if err != nil {
			return err
		}

		return next(context.ForSession(&res.SessionId), data)
	}
}

func communicateMiddleware(next Handler) Handler {
	return func(context *models.Context, data *models.Data) error {
		err := next(context, data)

		if err != nil {
			res := models.ErrorResponse{Error: "Something went wrong handling this request."}

			err = ws.Send(context, data.Route(), res)

			return err
		}

		return nil
	}
}

func Route(context *models.Context, data *models.Data) {
	route := data.Route().Value()

	log.Printf("Beginning Invoking '%s'", route)

	if handler, exists := handlers[route]; exists {
		err := handler(context, data)

		if err != nil {
			log.Printf("Error Invoking '%s': %s", route, err)
		}
	}
}