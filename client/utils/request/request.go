package request

import (
	"bytes"
	"client/config"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func CreateJSONRequest(method, endpoint string, data interface{}) *http.Request {
	url := strings.Join([]string{config.Envs.ServerHost, endpoint}, "")

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("converting to json failed: %v", err)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	return req
}
