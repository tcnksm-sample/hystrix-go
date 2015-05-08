package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", logger(HandleSubsystem))

	fmt.Println("==> Main server is started")
	log.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}

// HandleSubsystem send request to sub-system and extracts its response
func HandleSubsystem(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	resp, err := http.Get("http://localhost:9090")
	if err != nil {
		log.Println("failed to get response from sub-system:", err.Error())
		return
	}

	log.Println("success to get response from sub-system")

	w.WriteHeader(http.StatusOK)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Should not reach here
		panic(err)
	}
	w.Write(b)
}

// log is Handler wrapper function for logging
func logger(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		fn(w, r)
	}
}
