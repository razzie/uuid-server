package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/razzie/uuid-server/random"
	"github.com/razzie/uuid-server/random/seed"
)

func serveUUID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(uuid.New().String()))
}

func main() {
	port := flag.Int("port", 8080, "http port to listen on")
	apikey := flag.String("apikey", "", "random.org api key")
	flag.Parse()

	var seeder seed.Seeder
	if len(*apikey) > 0 {
		seeder = seed.RandomOrg(*apikey)
	} else {
		seeder = seed.RandomOrgHax()
	}

	seed := seed.RateLimit(seeder, 1, time.Minute)
	feed, _ := random.NewFeed(seed)
	uuid.SetRand(feed)

	http.HandleFunc("/", serveUUID)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
