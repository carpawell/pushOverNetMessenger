package api

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Server struct {
	Server *http.Server
	Router *httprouter.Router
}

type Opts struct {
	Port   string
	Routes Routes
}

func NewServer(opt *Opts) *Server {
	var srv = &Server{}

	srv.Router = httprouter.New()
	for _, route := range opt.Routes {
		srv.Router.Handle(route.Method, route.Path, route.Handler)
	}

	srv.Server = &http.Server{Addr: ":" + opt.Port, Handler: srv.Router}

	return srv
}

func (srv Server) Start() error {
	return srv.Server.ListenAndServe()
}

func (srv Server) Stop(ctx context.Context) error {
	return srv.Server.Shutdown(ctx)
}
