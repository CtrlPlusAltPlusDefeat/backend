package models

type Data struct {
	Message Wrapper
}

func NewData(request string) (*Data, error) {
	var message Wrapper
	err := message.Decode([]byte(request))

	if err != nil {
		return nil, err
	}

	return &Data{
		Message: message,
	}, nil
}
