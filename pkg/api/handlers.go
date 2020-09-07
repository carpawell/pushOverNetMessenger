package api

import (
	"encoding/json"
	"github.com/carpawell/pushOverNetMessenger/pkg/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (svc Service) GetMessagesCount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	keys, ok := r.URL.Query()["from"]
	if !ok || len(keys[0]) < 1 {
		errorResponse(w, FromParameterNotFound, http.StatusBadRequest)
		return
	}

	parsedTime, err := utils.ParseTime(keys[0])
	if err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	statistics := svc.Db.GetMessageStatistics(parsedTime)

	if err := json.NewEncoder(w).Encode(statistics); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
	}
}

func (svc Service) SendMessage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	message := &MessageToSend{}
	if err := json.NewDecoder(r.Body).Decode(message); err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	if message.Text == nil || *message.Text == "" {
		errorResponse(w, EmptyMessage, http.StatusBadRequest)
		return
	}

	response, err := svc.PushOverApp.SendMessage(*message.Text)
	if err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
	}

	svc.Db.AddNotification(*message.Text, response.Status)
}
