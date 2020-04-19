package main

import (
	"net/http"

	"github.com/google/uuid"
)

func serveUUID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(uuid.New().String()))
}

func main() {
	http.HandleFunc("/", serveUUID)
	http.ListenAndServe(":8080", nil)
}
