package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%d", rand.IntN(6)+1)

}

func main() {

	addr := ":8090"

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("error: %v", err)
	}

}
