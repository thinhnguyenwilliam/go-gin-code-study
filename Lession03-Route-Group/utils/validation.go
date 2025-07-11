package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

func ValidateSlug(fl validator.FieldLevel) bool {
	return slugRegex.MatchString(fl.Field().String())
}

func GetCustomErrorMessage(field, tag string, fe validator.FieldError) string {
	switch field {
	case "Lang":
		switch tag {
		case "required":
			return "Language is required"
		case "oneof":
			return fmt.Sprintf("Language must be one of: %s", fe.Param())
		}
	case "Slug":
		switch tag {
		case "required":
			return "Slug is required"
		case "min":
			return fmt.Sprintf("%s must be at least %s characters", field, fe.Param())
		case "max":
			return fmt.Sprintf("%s must be at most %s characters", field, fe.Param())
		case "slug":
			return "Slug can only contain lowercase letters, numbers, and hyphens"
		}

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
		case "min":
			return fmt.Sprintf("Search must be at least %s characters", fe.Param())
		case "max":
			return fmt.Sprintf("Search must be at most %s characters", fe.Param())
		case "email":
			return "Email must be a valid email address"
		}
	case "UUID":
		switch tag {
		case "uuid4":
			return "Invalid UUID format bro"
		}

	}
	return fmt.Sprintf("Invalid value for %s", field)
}

// utils/validator.go
func FormatValidationErrors(err error) map[string]string {
	formatted := map[string]string{}

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			field := fe.Field()
			tag := fe.Tag()
			formatted[strings.ToLower(field)] = GetCustomErrorMessage(field, tag, fe)
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
