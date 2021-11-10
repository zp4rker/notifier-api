package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/zp4rker/notifier-api/routes"
)

func main() {
	fmt.Println("Starting server...")
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("unable to start server")
	}
	fmt.Println("Server started!")

	fmt.Println("Setting up server...")
	routes.Init()
	go handleQuit(l)

	fmt.Println("Setup server!")
	http.Serve(l, nil)
}

func handleQuit(l net.Listener) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
	fmt.Println("\nStopping server...")
	l.Close()
	fmt.Println("Server stopped!")
}
