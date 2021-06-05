package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/technomunk/cookbook/recipe"
)

var db *sql.DB

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

func viewRecipeHandler(w http.ResponseWriter, r *http.Request) {
	rq := r.URL.Query()

	if rid, ok := parseInt64(rq.Get("rid")); ok {
		rcp, err := recipe.SearchById(db, rid)
		if err != nil {
			// TODO: figure out how to respond
			return
		}

		if rcp == nil {
			http.Error(w, "Recipe not found", http.StatusNotFound)
			return
		}

		err = recipeTemplates.ExecuteTemplate(w, "recipe.html", rcp.Recipe)
		if err != nil {
			http.Error(w, "Failed to populate template", http.StatusInternalServerError)
			return
		}
		return
	}

	// TODO: enumerate recipes on empty query
	http.Error(w, "Under development", http.StatusNotImplemented)
}

func editRecipeHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: shtuffs
	recipe := r.URL.Path[len("/food/add/"):]
	log.Println(recipe)
	fmt.Fprintf(w, "under development")
}

func addRecipeHandler(rw http.ResponseWriter, r *http.Request) {
	// TODO: require authentication
	switch r.Method {
	case http.MethodGet:
		http.ServeFile(rw, r, "content/food-add.html")

	case http.MethodPost:
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
