package api

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

type Service struct {
	Server *Server
	Router *httprouter.Router
}

func (svc Service) Start(ctx context.Context) error {
	svc.Router = httprouter.New()
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
		log.Fatalf("server Shutdown Failed:%s", err)
		return err
	}

	log.Printf("server stopped")
	return nil
}
