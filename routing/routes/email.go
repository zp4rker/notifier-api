package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type emailRequest struct {
	To      string
	From    string
	Sender  string
	Subject string
	Body    string
}

func EmailRoute(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		data, _ := json.Marshal(map[string]string{
			"error": "Error reading request body",
		})
		rw.Write(data)
		return
	}

	emailRequest := emailRequest{}
	if err := json.Unmarshal(body, &emailRequest); err != nil {
		data, _ := json.Marshal(map[string]interface{}{
			"error": "Error parsing request body",
			"request": map[string]string{
				"body": string(body),
			},
		})
		rw.Write(data)
		return
	}

	output := fmt.Sprintf("To: %v\n", emailRequest.To)
	output += fmt.Sprintf("From: %v <%v>\n", emailRequest.Sender, emailRequest.From)
	output += fmt.Sprintf("Subject: %v\n", emailRequest.Subject)
	output += fmt.Sprintf("\n%v\n", emailRequest.Body)

	rw.Write([]byte(output))
}
