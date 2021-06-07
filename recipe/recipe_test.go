package recipe_test

import (
	"testing"

	"github.com/technomunk/cookbook/recipe"
)

func TestSameAs(t *testing.T) {
	// Differentiate by product
	a := recipe.Recipe{"a", 1, []recipe.Ingredient{}, ""}
	b := recipe.Recipe{"b", 1, []recipe.Ingredient{}, ""}
	if a.SameAs(&b) {
		t.Fail()
	}

	// Differentiate by process
	a = recipe.Recipe{"a", 1, []recipe.Ingredient{}, "mix"}
	b = recipe.Recipe{"a", 1, []recipe.Ingredient{}, ""}
	if a.SameAs(&b) {
		t.Fail()
	}

	// Differentiate by rate
	a = recipe.Recipe{"a", 2, []recipe.Ingredient{}, ""}
	b = recipe.Recipe{"a", 1, []recipe.Ingredient{}, ""}
	if a.SameAs(&b) {
		t.Fail()
	}

	// Differentiate by ingredient count
	a = recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 1}, {"c", 1}}, ""}
	b = recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 1}}, ""}
	if a.SameAs(&b) {
		t.Fail()
	}

	// Differentiate by ingredient name
	a = recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 1}}, ""}
	b = recipe.Recipe{"a", 1, []recipe.Ingredient{{"c", 1}}, ""}
	if a.SameAs(&b) {
		t.Fail()
	}

	// Differentiate by ingredient rate
	a = recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 1}}, ""}
	b = recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 1}}, ""}
	if a.SameAs(&b) {
		t.Fail()
	}

	// Do not differentiate by address
	a = recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 1}}, "mix"}
	b = recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 1}}, "mix"}
	if !a.SameAs(&b) {
		t.Fail()
	}

	// Do not differentiate by ingredient order
	a = recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 1}, {"c", 1}}, "mix"}
	b = recipe.Recipe{"a", 1, []recipe.Ingredient{{"c", 1}, {"b", 1}}, "mix"}
	if !a.SameAs(&b) {
		t.Fail()
	}
}

// Check that cloning works as expected.
func TestClone(t *testing.T) {
	a := recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 1}, {"c", 1}}, "mix"}
	b := a.Clone()

	if !a.SameAs(&b) {
		t.Fail()
	}

	a.Ingredients = []recipe.Ingredient{}

	if a.SameAs(&b) {
		t.Fail()
	}
}

// Check that recipe.AdjustRate() functions as expected.
func TestAdjustRate(t *testing.T) {
	a := recipe.Recipe{"a", 1, []recipe.Ingredient{{"b", 2}}, ""}
	a.AdjustRate(2)

	if a.Rate != 2 {
		t.FailNow()
	}

	if a.Ingredients[0].Rate != 4 {
		t.Fail()
	}
}
