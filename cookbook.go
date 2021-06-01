package main

import (
	"fmt"
	"log"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello! The website is up!")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
