package recipe

import (
	"net/url"
	"strconv"
	"strings"
)

const (
	ErrCodeNoProduct = iota
	ErrCodeNoRate
	ErrCodeRateNotNumber
	ErrCodeNegativeRate
	ErrCodeIngredientCountMismatch
	ErrCodeNamelessIngredient
)

type ParseError struct {
	code int
}

func (e *ParseError) Error() string {
	switch e.code {
	case ErrCodeNoProduct:
		return "no product in provided form"

	case ErrCodeNoRate:
		return "no rate in provided form"

	case ErrCodeRateNotNumber:
		return "rate provided is not a valid number"

	case ErrCodeNegativeRate:
		return "rate must be a positive number"

	case ErrCodeIngredientCountMismatch:
		return "ingredient name and rate count must be the same"

	case ErrCodeNamelessIngredient:
		return "ingredient is missing a name"

	default:
		return "unknown error"
	}
}

func (e *ParseError) Code() int {
	return e.code
}

// Build a recipe instance using the values provided in a url form.
func ParseRecipe(form url.Values) (*Recipe, error) {
	product, err := parseProduct(form)
	if err != nil {
		return nil, err
	}

	rate, err := parseRate(form)
	if err != nil {
		return nil, err
	}

	process, err := parseProcess(form)
	if err != nil {
		return nil, err
	}

	ingredients, err := parseIngredients(form)
	if err != nil {
		return nil, err
	}

	return &Recipe{product, rate, ingredients, process}, nil
}

func parseProduct(form url.Values) (string, error) {
	result := strings.ToLower(strings.TrimSpace(form.Get("product")))
	if result == "" {
		return "", &ParseError{ErrCodeNoProduct}
	}
	return result, nil
}

func parseRate(form url.Values) (Rate, error) {
	str := strings.TrimSpace(form.Get("rate"))
	if str == "" {
		return 0, &ParseError{ErrCodeNoRate}
	}

	rate, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, &ParseError{ErrCodeRateNotNumber}
	}

	if rate <= 0 {
		return 0, &ParseError{ErrCodeNegativeRate}
	}

	return Rate(rate), nil
}

func parseProcess(form url.Values) (string, error) {
	return strings.ToLower(strings.TrimSpace(form.Get("process"))), nil
}

func parseIngredients(form url.Values) ([]Ingredient, error) {
	names := form["ingredient-name"]
	rates := form["ingredient-rate"]

	if len(names) != len(rates) {
		return nil, &ParseError{ErrCodeIngredientCountMismatch}
	}

	result := make([]Ingredient, 0, len(names))

	for i := range names {
		// The form is allowed to have empty indices for usability reasons
		if names[i] == "" && rates[i] == "" {
			continue
		}

		names[i] = strings.ToLower(strings.TrimSpace(names[i]))
		if names[i] == "" {
			return nil, &ParseError{ErrCodeNamelessIngredient}
		}

		rate, err := strconv.ParseFloat(strings.TrimSpace(rates[i]), 64)
		if err != nil {
			return nil, &ParseError{ErrCodeRateNotNumber}
		}

		if rate <= 0 {
			return nil, &ParseError{ErrCodeNegativeRate}
		}

		result = append(result, Ingredient{names[i], Rate(rate)})
	}

	return result, nil
}
