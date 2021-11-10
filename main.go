package main

import (
	"fmt"
	"net/http"

	"github.com/zp4rker/notifier-api/routes"
)

func main() {
	fmt.Println("Started server!")
	http.HandleFunc("/", routes.NotFoundRoute)
	http.HandleFunc("/hello-world", routes.HelloWorldRoute)
	http.HandleFunc("/json-test", routes.JsonTestRoute)
	http.ListenAndServe(":8080", nil)
}
