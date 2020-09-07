package api

type MessageToSend struct {
	Text *string `json:"message"`
}

type Error struct {
	ErrorMessage string `json:"error"`
}
