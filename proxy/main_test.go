package main

import (
	"io"
	"net/http"
	"testing"
)

func TestMainFunc(t *testing.T) {
	go main()

	req, _ := http.NewRequest("GET", "http://localhost:8080/api", nil)
	resp, _ := http.DefaultClient.Do(req)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status code is not 200: %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "Hello from API" {
		t.Errorf("response is not \"Hello from API\": %s", string(body))
	}

	req, _ = http.NewRequest("GET", "http://localhost:8080/", nil)
	resp, _ = http.DefaultClient.Do(req)

	body, _ = io.ReadAll(resp.Body)
	if string(body) == "Hello from API" {
		t.Errorf("response is \"Hello from API\": %s", string(body))
	}
}
