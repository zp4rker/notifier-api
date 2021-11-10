package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func JsonTestRoute(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		data, _ := json.Marshal(map[string]string{
			"error": "Error reading request body",
		})
		rw.Write(data)
		return
	}

	inputMap := make(map[string]interface{})
	if err := json.Unmarshal(body, &inputMap); err != nil {
		data, _ := json.Marshal(map[string]string{
			"error": "Error parsing request body",
		})
		rw.Write(data)
		return
	}

	output, _ := json.Marshal(map[string]string{
		"message": "Successfully received JSON!",
	})
	rw.Write(output)
}
