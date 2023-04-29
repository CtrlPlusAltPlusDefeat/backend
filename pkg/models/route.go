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
