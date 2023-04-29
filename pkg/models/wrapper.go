package models

type Wrapper struct {
	Service string `json:"service"`
	Action  string `json:"action"`
	Data    string `json:"data"`
}
