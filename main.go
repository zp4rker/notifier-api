package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/zp4rker/notifier-api/routing"
)

func main() {
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("unable to start server")
	}
	fmt.Println("Server started!")

	fmt.Println("Setting up server...")
	routing.Init()
	go handleQuit(l)
	fmt.Println("Setup server!")

	if err := http.Serve(l, nil); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed), errors.Is(err, net.ErrClosed):
			fmt.Println("Server stopped!")
		default:
			panic(err)
		}
	}
}

func handleQuit(l net.Listener) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	fmt.Println("\nStopping server...")
	l.Close()
	fmt.Println("Server stopped!")
}
