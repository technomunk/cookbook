package recipe_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/technomunk/cookbook/recipe"
)

// Check that recipe.FindRecipe() returns newly inserted recipe.
func TestInsertAndFindRecipe(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if err := recipe.CreateTables(db); err != nil {
		t.Fatal(err)
	}

	recipes, err := recipe.Find(db, "dough")
	if err != nil {
		t.Fatal(err)
	}

	if len(recipes) != 0 {
		t.Fatal("recipe was already present in a database")
	}

	_, err = recipe.Insert(db, &recipe.ExampleRecipe)
	if err != nil {
		t.Fatal(err)
	}

	recipes, err = recipe.Find(db, "dough")
	if err != nil {
		t.Fatal(err)
	}

	if len(recipes) != 1 || recipe.ExampleRecipe.SameAs(&recipes[0].Recipe) {
		t.Fatal("did not find the expected recipes")
	}
}
