package models

type Route struct {
	service *string
	action  *string
}

func NewRoute(service *string, action *string) *Route {
	return &Route{
		service: service,
		action:  action,
	}
}

func (a *Route) Service() *string {
	return a.service
}

func (a *Route) Action() *string {
	return a.action
}

func (a *Route) Value() string {
	return *a.service + "|" + *a.action
}

func PlayerLeave() *Route {
	service := "lobby"
	action := "player-left"
	return NewRoute(&service, &action)
}

func PlayerJoined() *Route {
	service := "lobby"
	action := "player-joined"
	return NewRoute(&service, &action)
}

func JoinedLobby() *Route {
	service := "lobby"
	action := "join"
	return NewRoute(&service, &action)
}

func SetSession() *Route {
	service := "player"
	action := "set-session"
	return NewRoute(&service, &action)

}
