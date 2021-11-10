package routing

import (
	"net/http"

	"github.com/zp4rker/notifier-api/routing/routes"
)

func Init() {
	http.HandleFunc("/", routes.NotFoundRoute)
	http.HandleFunc("/hello-world", routes.HelloWorldRoute)
	http.HandleFunc("/json-test", routes.JsonTestRoute)
}
