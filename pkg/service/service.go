package service

import (
	"context"
	"github.com/carpawell/pushOverNetMessenger/pkg/api"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type Service struct {
	Server *http.Server
	Router *httprouter.Router
}

func (svc Service) Start(ctx context.Context) error {
	svc.Router = httprouter.New()
	svc.Server = &http.Server{Addr: ":8080", Handler: svc.Router}

	for _, route := range api.GETRoutes {
		svc.Router.GET(route.Path, route.Handler)
	}
	for _, route := range api.POSTRoutes {
		svc.Router.POST(route.Path, route.Handler)
	}

	go func() {
		if err := svc.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := svc.Server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
		return err
	}

	log.Printf("server stopped")
	return nil
}
