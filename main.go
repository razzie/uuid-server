package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/razzie/random"
	"github.com/razzie/random/seed"
)

func serveUUID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(uuid.New().String()))
	logRequest(r)
}

func logRequest(r *http.Request) {
	ip := r.Header.Get("X-REAL-IP")
	if len(ip) == 0 {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	log.Println(ip, r.RequestURI)
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
