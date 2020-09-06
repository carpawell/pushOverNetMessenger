package api

type MessageToSend struct {
	Message *string `json:"message"`
}

type Error struct {
	ErrorMessage string `json:"error"`
}
