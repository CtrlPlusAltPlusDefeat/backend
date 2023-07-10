package models

import (
	"encoding/json"
	"log"
)

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
	err := json.Unmarshal(d.data, req)
	if err != nil {
		log.Printf("Failed to decode data: %s", d.data)
	}
	return err
}

func (d *Data) Data() string {
	return string(d.data)
}

func (d *Data) Route() *Route {
	return d.route
}
