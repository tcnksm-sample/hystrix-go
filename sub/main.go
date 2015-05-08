package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", logger(HandleHeavyJob))

	fmt.Println("==> Sub server is started")
	log.Println("listening on :9090")
	http.ListenAndServe(":9090", nil)
}

// HandleHeavyJob send request to sub-system and extracts its response
func HandleHeavyJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// This takes time and sometimes fail...
	time.Sleep(1 * time.Second)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}

// log is Handler wrapper function for logging
func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		fn(w, r)
	}
}
