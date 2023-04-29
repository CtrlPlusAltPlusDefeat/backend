package models

import "encoding/json"

type Data struct {
	Action *Route
	value  []byte
}

func NewData(request string) (*Data, error) {
	message := Wrapper{}
	err := json.Unmarshal([]byte(request), &message)

	if err != nil {
		return nil, err
	}

	bytes := []byte(message.Data)

	return &Data{
		Action: &Route{
			action:  &message.Action,
			service: &message.Service,
		},
		value: bytes,
	}, nil
}

func (d *Data) DecodeTo(req interface{}) error {
	return json.Unmarshal(d.value, req)
}
