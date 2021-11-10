package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/zp4rker/notifier-api/routes"
)

func main() {
	srv := &http.Server{Addr: ":8080"}
	routes.Init()
	fmt.Println("Started server!")
	go handleQuit(srv)
	srv.ListenAndServe()
}

func handleQuit(srv *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	fmt.Println("\nStopping server...")
	srv.Shutdown(context.TODO())
}
