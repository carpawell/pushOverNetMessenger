package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (svc Service) GetMessages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "test")
}

func (svc Service) SendMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "test")
}
