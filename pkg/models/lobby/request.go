package lobby

type CreateAndJoinRequest struct {
	Name string `json:"name"`
}

type SetNameRequest struct {
	Text string `json:"text"`
}
