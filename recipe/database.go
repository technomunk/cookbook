package recipe

import (
	"database/sql"
	"strconv"
)

// A database entry with a provided recipe
type RecipeEntry struct {
	Id     int64
	Recipe Recipe
}

// Create tables necessary to store recipe and their ingredients.
func CreateTables(db *sql.DB) error {
	const createRecipesTableQuery = `CREATE TABLE recipe(
"recipeid" INTEGER PRIMARY KEY,
"product" TEXT NOT NULL,
"rate" REAL NOT NULL,
"process" TEXT);
CREATE INDEX idx_product ON recipe(product);
`
	const createIngredientsTableQuery = `CREATE TABLE ingredient(
"recipeid" INTEGER NOT NULL,
"name" TEXT NOT NULL,
"rate" REAL NOT NULL,
FOREIGN KEY(recipeid) REFERENCES recipe(recipeid)
);
CREATE INDEX idx_recipeid ON ingredient(recipeid);
CREATE INDEX idx_name ON ingredient(name);`

	_, err := db.Exec(createRecipesTableQuery)
	if err != nil {
		return err
	}

	_, err = db.Exec(createIngredientsTableQuery)
	return err
}

// Insert a new recipe into the database.
func Insert(db *sql.DB, r *Recipe) (int64, error) {
	const insertRecipeQuery = `INSERT INTO recipe(product, rate, process) VALUES(?, ?, ?);`
	const insertIngredientQuery = `INSERT INTO ingredient(recipeid, name, rate) VALUES (?, ?, ?);`

	// Initialize the transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback() // the rollback is ignored if the transaction was committed.

	// Insert the recipe without ingredients
	result, err := tx.Exec(insertRecipeQuery, r.Product, r.Rate, r.Process)
	if err != nil {
		return 0, err
	}

	// Prepare ingredient query
	rid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare(insertIngredientQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close() // prepared statements should be closed after use

	// Execute ingredient query for all the recipe ingredients
	for _, ingredient := range r.Ingredients {
		_, err = stmt.Exec(rid, ingredient.Name, ingredient.Rate)
		if err != nil {
			return 0, err
		}
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return rid, nil
}

// Select recipes that produce provided product without ingredients.
func selectRecipes(db *sql.DB, product string) ([]RecipeEntry, error) {
	rows, err := db.Query(`SELECT * FROM recipe WHERE product=?;`, product)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	recipes := make([]RecipeEntry, 0)

	for rows.Next() {
		var (
			recipeId int64
			product  string
			rate     float64
			process  string
		)
		err = rows.Scan(&recipeId, &product, &rate, &process)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, RecipeEntry{recipeId, Recipe{Product: product, Rate: Rate(rate), Process: process}})
	}

	return recipes, nil
}

// Gather recipe ingredients for each of the recipes in the provided slice.
func selectRecipeIngredients(db *sql.DB, recipes []RecipeEntry) error {
	if len(recipes) == 0 {
		return nil
	}

	// Collect recipe ids
	rids := ""
	for i := range recipes {
		rids += strconv.Itoa(i)
		if i+1 < len(recipes) {
			rids += ","
		}
	}

	rows, err := db.Query(`SELECT * FROM ingredient WHERE recipeid IN(?) ORDER BY recipeid;`, rids)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rcp := &recipes[0]; rows.Next(); {
		var (
			recipeId int64
			name     string
			rate     float64
		)
		err = rows.Scan(&recipeId, &name, &rate)
		if err != nil {
			return err
		}

		if rcp.Id != recipeId {
			for idx := range recipes {
				if recipeId == recipes[idx].Id {
					rcp = &recipes[idx]
					break
				}
			}
		}

		rcp.Recipe.Ingredients = append(rcp.Recipe.Ingredients, Ingredient{name, Rate(rate)})
	}

	return nil
}

// Search the databases for recipes that create provided product.
func Find(db *sql.DB, product string) ([]RecipeEntry, error) {
	recipes, err := selectRecipes(db, product)
	if err != nil {
		return nil, err
	}
	err = selectRecipeIngredients(db, recipes)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}
