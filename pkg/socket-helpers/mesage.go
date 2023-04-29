package socket_helpers

type Message struct {
	Service string `json:"service"`
	Action  string `json:"action"`
	Data    string `json:"data"`
}
