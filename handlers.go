package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/technomunk/cookbook/recipe"
)

var db *sql.DB

// Base catch-all handler for requests not covered by other handlers
func rootHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(rw, r)
		return
	}

	user, _, ok := r.BasicAuth()
	if !ok {
		http.Error(rw, "Not authenticated", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(rw, "Hello %s! The website is up!", user)
}

// Provides an overview of available recipes or if a valid query is provided a view of a particular recipe.
func viewRecipeHandler(rw http.ResponseWriter, r *http.Request) {
	rq := r.URL.Query()

	if rid, ok := parseInt64(rq.Get("rid")); ok {
		rcp, err := recipe.SearchById(db, rid)
		if err != nil {
			log.Println(err)
			http.Error(rw, "Error searching for recipe", http.StatusInternalServerError)
			return
		}

		if rcp == nil {
			http.Error(rw, "Recipe not found", http.StatusNotFound)
			return
		}

		err = recipeTemplates.ExecuteTemplate(rw, "recipe.html", rcp.Recipe)
		if err != nil {
			log.Println(err)
			http.Error(rw, "Failed to populate template", http.StatusInternalServerError)
			return
		}
		return
	}

	// TODO: handle other queries

	rcps, err := recipe.EnumerateAll(db)
	if err != nil {
		log.Println(err)
		http.Error(rw, "Failed to get recipes", http.StatusInternalServerError)
		return
	}

	recipeTemplates.ExecuteTemplate(rw, "overview.html", rcps)
}

func editRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: shtuffs
	recipe := r.URL.Path[len("/food/add/"):]
	log.Println(recipe)
	fmt.Fprintf(w, "under development")
}

// Provides the recipe creation form or processing incoming POST requests.
func addRecipeHandler(rw http.ResponseWriter, r *http.Request) {
	// TODO: require authentication
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(rw, r, "content/food-add.html")

	case http.MethodPost:
		// TODO: add JSON content handler
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parcing form:", err)
			http.Error(rw, "Invalid form", http.StatusBadRequest)
			return
		}

		rcp, err := recipe.ParseRecipe(r.PostForm)
		if err != nil {
			log.Println("Failed to parse recipe:", err, r.PostForm)
			http.Error(rw, "Invalid form", http.StatusBadRequest)
			return
		}

		rid, err := recipe.Insert(db, rcp)
		if err != nil {
			log.Println("Failed to submit recipe:", err)
			http.Error(rw, "Internal server error", http.StatusInternalServerError)
			return
		}

		http.Redirect(rw, r, fmt.Sprintf("/food/view/?rid=%d", rid), http.StatusSeeOther)
	}
}

// Handler that fetches allowed content from /content/ folder.
func contentHandler(rw http.ResponseWriter, r *http.Request) {
	const contentPrefixLen = len("/content/")
	item := r.URL.Path[contentPrefixLen:]

	if item == "ingredient-input.js" {
		http.ServeFile(rw, r, "content/ingredient-input.js")
		return
	}

	http.Error(rw, "Content not found", http.StatusNotFound)
}
