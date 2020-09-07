package api

import (
	"context"
	"github.com/carpawell/pushOverNetMessenger/pkg/constants"
	"github.com/carpawell/pushOverNetMessenger/pkg/pushApp"
	"github.com/carpawell/pushOverNetMessenger/pkg/storage"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type Service struct {
	Server      *Server
	Router      *httprouter.Router
	PushOverApp *pushApp.PushApp
	Db          *storage.Storage
}

func (svc Service) Start(ctx context.Context) error {
	// Connection to database
	db, err := storage.New()
	if err != nil {
		return err
	}

	// Initializing service
	svc.Db = db
	svc.Router = httprouter.New()
	svc.PushOverApp = pushApp.New(constants.App, constants.User)
	opt := &Opts{Port: "8080", Routes: svc.routes()}
	svc.Server = NewServer(opt)

	go func() {
		if err := svc.Server.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()

	log.Printf("server started")
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := svc.Server.Stop(ctxShutDown); err != nil {
		log.Fatalf("server shutdown failed: %s", err)
		return err
	}

	log.Printf("server stopped")
	return nil
}
