package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Known user credentials. Loaded from secret file at program start.
var users map[string]string

// Load user credentials from secret file.
func init() {
	// TODO: store user data more securely
	bytes, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(bytes, &users)
	if err != nil {
		log.Fatal(err)
	}
}

// Require the request to be properly authenticated.
//
// Automatically responds with appropriate errors.
// Returns true if the authentication was successful.
func requireAuth(rw http.ResponseWriter, r *http.Request) bool {
	un, pwd, ok := r.BasicAuth()

	if !ok {
		rw.Header().Set("WWW-Authenticate", "basic")
		http.Error(rw, "Not authenticated", http.StatusUnauthorized)
		return false
	}

	cpwd, ok := users[un]
	if !ok || cpwd != pwd {
		rw.Header().Set("WWW-Authenticate", "basic")
		http.Error(rw, "Invalid username or password", http.StatusUnauthorized)
		return false
	}

	return true
}
