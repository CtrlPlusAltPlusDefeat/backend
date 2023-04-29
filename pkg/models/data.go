package models

import "encoding/json"

type Data struct {
	route *Route
	data  []byte
}

func NewData(request string) (*Data, error) {
	message := Wrapper{}
	err := json.Unmarshal([]byte(request), &message)

	if err != nil {
		return nil, err
	}

	data := []byte(message.Data)

	return &Data{
		route: &Route{
			action:  &message.Action,
			service: &message.Service,
		},
		data: data,
	}, nil
}

func (d *Data) DecodeTo(req interface{}) error {
	return json.Unmarshal(d.data, req)
}

func (d *Data) Route() *Route {
	return d.route
}
