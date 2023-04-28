package models

type Connection struct {
	ConnectionId string `dynamodbav:"ConnectionId"`
	SessionId    string `dynamodbav:"SessionId"`
}
