package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (svc Service) GetMessages(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (svc Service) SendMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	message := &MessageToSend{}
	if err := json.NewDecoder(r.Body).Decode(message); err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	if message.Message == nil || *message.Message == "" {
		errorResponse(w, EmptyMessageError, http.StatusBadRequest)
		return
	}

	response, err := svc.PushOverApp.SendMessage(*message.Message)
	if err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
	}
}
