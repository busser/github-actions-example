package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/busser/github-actions-example/foobar"
)

const addr = ":8080"

func main() {
	http.HandleFunc("/", foobarHandler)
	log.Printf("Listening for requests on %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func foobarHandler(w http.ResponseWriter, r *http.Request) {
	lengthParam := r.URL.Query().Get("length")
	if len(lengthParam) == 0 {
		http.Error(w, "missing parameter: length", http.StatusBadRequest)
		return
	}

	length, err := strconv.Atoi(lengthParam)
	if err != nil {
		http.Error(w, "invalid length", http.StatusBadRequest)
		return
	}

	seq, err := foobar.Sequence(length)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to compute sequence: %s", err.Error()), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", seq)
}
