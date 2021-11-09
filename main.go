package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zp4rker/notifier-api/routes"
)

func main() {
	fmt.Println("Started server!")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(map[string]interface{}{
			"message": "Hello, World!",
		})
		if err != nil {
			rw.Write([]byte(err.Error()))
		} else {
			rw.Write(data)
		}
	})
	http.HandleFunc("/hello-world", routes.HelloWorldRoute)
	http.ListenAndServe(":8080", nil)
}
