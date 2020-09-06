package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Route struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request, httprouter.Params)
}

var GETRoutes = []Route{
	{"/messages", GetMessages},
}

var POSTRoutes = []Route{
	{"/messages", SendMessage},
}
