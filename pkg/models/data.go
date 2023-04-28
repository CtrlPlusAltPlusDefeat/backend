package models

type Data struct {
	Message Wrapper
}

func NewData(message Wrapper) *Data {
	return &Data{
		Message: message,
	}
}
