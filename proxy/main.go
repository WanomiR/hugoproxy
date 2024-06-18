package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	proxy := NewReverseProxy("hugo", "1313")
	r.Use(proxy.ReverseProxy)

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello from API"))
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
