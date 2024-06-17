package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	r := chi.NewRouter()

	// ...

	http.ListenAndServe(":8080", r)
}

const content = ``

func WorkerTest() {
	t := time.NewTicker(1 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := os.WriteFile("/app/static/_index.md", []byte(fmt.Sprintf(content, b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
