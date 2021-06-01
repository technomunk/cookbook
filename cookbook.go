package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/technomunk/cookbook/recipe"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello! The website is up!")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: the exact recipe should be gotten from url
	err := recipe.Templates.ExecuteTemplate(w, "recipe.txt", recipe.ExampleRecipe)
	if err != nil {
		http.Error(w, "Failed to populate template", http.StatusInternalServerError)
	}
}

func main() {
	port := flag.Int("port", 8080, "provide the port to bind to")
	flag.Parse()

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/view/", viewHandler)

	log.Println("Listening on", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
