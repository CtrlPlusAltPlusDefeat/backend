package models

import "encoding/json"

type Data struct {
	route *Route
	value []byte
}

func NewData(request string) (*Data, error) {
	message := Wrapper{}
	err := json.Unmarshal([]byte(request), &message)

	if err != nil {
		return nil, err
	}

	bytes := []byte(message.Data)

	return &Data{
		route: &Route{
			action:  &message.Action,
			service: &message.Service,
		},
		value: bytes,
	}, nil
}

func (d *Data) DecodeTo(req interface{}) error {
	return json.Unmarshal(d.value, req)
}

func (d *Data) Route() *Route {
	return d.route
}
