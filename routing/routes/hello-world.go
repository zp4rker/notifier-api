package routes

import (
	"encoding/json"
	"net/http"
)

func HelloWorldRoute(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	data, _ := json.Marshal(map[string]string{"hello": "world"})
	rw.Write(data)
}
