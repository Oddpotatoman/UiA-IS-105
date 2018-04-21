package main

import (
	"net/http"
	"fmt"
)

func main () {
	http.HandleFunc("/", client)
	http.ListenAndServe(":8080", nil)
}
func client (w http.ResponseWriter, r * http.Request) {
	fmt.Fprint(w, "Hello client")
}