package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler httprouter.Handle
}

type Routes = []Route

func (svc Service) routes() []Route {
	return Routes{
		{
			Method:  http.MethodGet,
			Path:    "/messages/statistics",
			Handler: svc.GetMessagesCount,
		},
		{
			Method:  http.MethodPost,
			Path:    "/messages",
			Handler: svc.SendMessage,
		},
	}
}
