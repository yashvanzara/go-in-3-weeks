package main

import (
	"log"
	"net/http"
)

func main() {
	ph := newProverbsHandler()

	mux := http.NewServeMux()
	mux.Handle("GET /proverbs/", ph)
	mux.Handle("GET /proverbs/{id}", ph)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}
