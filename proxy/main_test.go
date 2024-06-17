package main

import (
	"io"
	"net/http"
	"testing"
)

func TestMainFunc(t *testing.T) {
	go main()

	req, _ := http.NewRequest("GET", "http://localhost:8080/api", nil)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		t.Errorf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code is not 200: %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "Hello from API" {
		t.Errorf("response is not \"Hello from API\": %s", string(body))
	}
}
