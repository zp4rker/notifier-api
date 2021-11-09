package routes

import (
	"encoding/json"
	"net/http"
)

func HelloWorldRoute(rw http.ResponseWriter, req *http.Request) {
	data, _ := json.Marshal(map[string]string{"hello": "world"})
	rw.Write(data)
}
