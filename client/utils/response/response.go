package response

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetErrorFromResponse(response *http.Response) error {
	body := readResponseBody(response)

	var resp map[string]string
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Fatalf("converting from JSON failed: %v", err)
	}

	return fmt.Errorf(resp["error"])
}

func UnmarshalResponse(response *http.Response) map[string]string {
	body := readResponseBody(response)

	var resp map[string]string
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Fatalf("converting from JSON failed: %v", err)
	}

	return resp
}

func readResponseBody(response *http.Response) []byte {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error read request body: %v", err)
	}

	return body
}
