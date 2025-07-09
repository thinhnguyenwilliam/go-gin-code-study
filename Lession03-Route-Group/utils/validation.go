package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var alphaNumSpaceRegex = regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)

func AlphaNumSpace(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	return alphaNumSpaceRegex.MatchString(val)
}

// ValidateSearch checks if a search string is valid
func ValidateSearch(search string) error {
	search = strings.TrimSpace(search)

	if len(search) < 3 || len(search) > 50 {
		return errors.New("search must be between 3 and 50 characters")
	}

	// Only allow letters, numbers, and spaces
	validSearch := regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)
	if !validSearch.MatchString(search) {
		return errors.New("search may only contain letters, numbers, and spaces")
	}

	return nil
}

// ValidateLimit parses and validates the limit string
func ValidateLimit(limitStr string) (int, error) {
	if limitStr == "" {
		return 10, nil // default limit
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		return 0, errors.New("limit must be a positive number")
	}

	return limit, nil
}
