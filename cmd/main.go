package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var sleepTime time.Duration

// create http server listening on port 9090
func main() {
	// get the time to sleep from environment variable
	s := os.Getenv("SLEEP_TIME")
	if s == "" {
		s = "1s"
	}
	// convert string to time.Duration
	sleepTime, _ = time.ParseDuration(s)

	// create http server listening on port 9090
	fmt.Println("Starting server on port 9090")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9090", nil)
}

// handler waits for a request and returns a response
func handler(w http.ResponseWriter, r *http.Request) {
	// write response
	fmt.Printf("Received request from %s\n", r.RemoteAddr)
	time.Sleep(sleepTime)
	w.Write([]byte("Hello World\n"))
}