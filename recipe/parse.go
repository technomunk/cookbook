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

	default:
		return "unknown error"
	}
}

func (e *ParseError) Code() int {
	return e.code
}

func ParseRecipe(form url.Values) (*Recipe, error) {
	product := strings.TrimSpace(form.Get("product"))
	if product == "" {
		return nil, &ParseError{ErrCodeNoProduct}
	}

	rateStr := strings.TrimSpace(form.Get("rate"))
	if rateStr == "" {
		return nil, &ParseError{ErrCodeNoRate}
	}

	rate, err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		return nil, &ParseError{ErrCodeRateNotNumber}
	}

	if rate <= 0 {
		return nil, &ParseError{ErrCodeNegativeRate}
	}

	process := strings.TrimSpace(form.Get("process"))

	// TODO: parse ingredients
	return &Recipe{Product: product, Rate: Rate(rate), Process: process}, nil
}
