package main

import (
	"context"
	"github.com/carpawell/pushOverNetMessenger/pkg/api"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Main service
	svc := api.Service{}

	// Chanel for OS interruptions
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call: %v", oscall)
		cancel()
	}()

	if err := svc.Start(ctx); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
