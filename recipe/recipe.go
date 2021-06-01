package recipe

import (
	"fmt"
	"text/template"
)

// Relative amount of how much product is made or required by a recipe.
type Rate float32

// A single requirement for a recipe.
type Ingredient struct {
	Name string
	Rate Rate
}

// A blueprint how to create something.
type Recipe struct {
	// The result of the recipe
	Product string
	// The amount produced by the recipe
	Rate Rate
	// Requirements for the recipe
	Ingredients []Ingredient
	// The process by which the product is made
	Process string
}

func (i *Ingredient) String() string {
	return fmt.Sprintf("%fx\"%s\"", float32(i.Rate), i.Name)
}

var Templates *template.Template
var ExampleRecipe = Recipe{"dough", 1.63, []Ingredient{{"flour", 1}, {"water", .6}, {"salt", .02}, {"yeast", .01}}, "mix"}

func init() {
	Templates = template.Must(template.ParseFiles("content/tmpl/recipe.txt"))
}
