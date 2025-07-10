package utils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GetCustomErrorMessage(field, tag string) string {
	switch field {
	case "ID":
		switch tag {
		case "gt":
			return "ID must be greater than 0"
		}
	case "Search":
		switch tag {
		case "required":
			return "Search is required"
		case "alphanumspace":
			return "Search can only contain letters, numbers, and spaces"
		}
	case "UUID":
		switch tag {
		case "uuid4":
			return "Invalid UUID format bro"
		}

	}
	return "Invalid value"
}

// utils/validator.go
func FormatValidationErrors(err error) map[string]string {
	formatted := map[string]string{}

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := fe.Field()
			tag := fe.Tag()
			formatted[field] = GetCustomErrorMessage(field, tag)
		}
	}

	return formatted
}

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
