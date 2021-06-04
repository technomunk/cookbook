package main

import (
	"log"
	"net/http"
)

func logged(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		h(w, r)
	}
}
