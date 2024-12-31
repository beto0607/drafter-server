package main

import (
	"context"
	"drafter/src/core"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	core.ConnectToDB()

	server := core.InitServer()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Println("Starting server on " + server.Addr)
		server.ListenAndServe()
	}()

	<-ctx.Done()

	server.Shutdown(context.TODO())
	core.DisconnectDB()
	log.Println("Server shutdown")
	log.Println("final")
}
