package main

import (
	"fmt"
	"log"
	"net/http"

)

func main() {
	http.HandleFunc("/", juni)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func juni(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client")
}