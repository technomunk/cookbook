package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var recipeTemplates *template.Template

func init() {
	recipeTemplates = template.Must(template.ParseFiles("content/tmpl/recipe.txt"))
	var err error
	db, err = sql.Open("sqlite3", "food.db")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	port := flag.Int("port", 8080, "provide the port to bind to")
	flag.Parse()

	http.HandleFunc("/", logged(rootHandler))
	http.HandleFunc("/food/view/", logged(viewRecipeHandler))
	http.HandleFunc("/food/add/", logged(addRecipeHandler))

	log.Println("Listening on", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
