package routes

import "net/http"

func NotFoundRoute(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte("404 - Page not found"))
}
