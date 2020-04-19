package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/razzie/uuid-server/random"
)

func serveUUID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(uuid.New().String()))
}

func main() {
	seed := make(chan uint64, 1)
	seed <- 0
	feed, _ := random.NewFeed(seed)
	uuid.SetRand(feed)

	http.HandleFunc("/", serveUUID)
	http.ListenAndServe(":8080", nil)
}
