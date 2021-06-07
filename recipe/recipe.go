package recipe

import (
	"fmt"
)

// Relative amount of how much product is made or required by a recipe.
type Rate float64

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
	return fmt.Sprintf("%fx\"%s\"", float64(i.Rate), i.Name)
}

// Check whether an ingredient is contained within ingredients slice.
func (i *Ingredient) IsIn(slice []Ingredient) bool {
	for idx := range slice {
		if slice[idx] == *i {
			return true
		}
	}
	return false
}

// Check if one recipe is exactly the same as another one.
//
// Has O(N^2) where N is the number of ingredients of the recipe.
func (a *Recipe) SameAs(b *Recipe) bool {
	if len(a.Ingredients) != len(b.Ingredients) || a.Product != b.Product || a.Process != b.Process || a.Rate != b.Rate {
		return false
	}

	for i := range a.Ingredients {
		if !a.Ingredients[i].IsIn(b.Ingredients) {
			return false
		}
	}

	return true
}

// Scale the rates of the recipe and ingredients to achieve provided result rate.
func (r *Recipe) AdjustRate(rate Rate) {
	m := rate / r.Rate
	r.Rate *= m
	for i := range r.Ingredients {
		r.Ingredients[i].Rate *= m
	}
}

// Perform a deep copy of the recipe.
func (r *Recipe) Clone() Recipe {
	ingrs := make([]Ingredient, len(r.Ingredients))
	for i := range r.Ingredients {
		ingrs[i] = r.Ingredients[i]
	}

	return Recipe{r.Product, r.Rate, ingrs, r.Process}
}

var ExampleRecipe = Recipe{"dough", 1.63, []Ingredient{{"flour", 1}, {"water", .6}, {"salt", .02}, {"yeast", .01}}, "mix"}
