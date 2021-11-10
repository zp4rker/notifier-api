package routes

import "net/http"

func Init() {
	http.HandleFunc("/", NotFoundRoute)
	http.HandleFunc("/hello-world", HelloWorldRoute)
	http.HandleFunc("/json-test", JsonTestRoute)
}
