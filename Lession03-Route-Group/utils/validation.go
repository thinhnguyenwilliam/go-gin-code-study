package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Accept only .jpg, .jpeg, .png
func ValidateImageExtension(fl validator.FieldLevel) bool {
	url := fl.Field().String()

	log.Printf("[Validation] Checking image URL: %s", url)

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}

	for _, ext := range allowedExtensions {
		if strings.HasSuffix(strings.ToLower(url), ext) {
			log.Printf("[Validation] Passed for extension: %s", ext)
			return true
		}
	}

	log.Println("[Validation] Failed: unsupported image extension")
	return false
}

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

func ValidateSlug(fl validator.FieldLevel) bool {
	return slugRegex.MatchString(fl.Field().String())
}

// This is a nested map (map of maps)
// msg := staticMessages["Name"]["required"]
// fmt.Println(msg) // Output: Name is required
var staticMessages = map[string]map[string]string{
	"URL": {
		"required": "URL is required bro",
		"imgext":   "Image URL must end with .jpg, .jpeg, or .png",
	},
	"Name": {
		"required": "Name is required",
		"min":      "Name Product must be at least 3 characters",
		"max":      "Name must be at most 100 characters",
	},
	"Email": {
		"email": "Email must be a valid email address",
	},
	"Price": {
		"gt": "Price must be greater than 0",
	},
	"Stock": {
		"gte": "Stock must be greater than or equal to 0",
	},
	"Lang": {
		"required": "Language is required",
	},
	"Slug": {
		"required": "Slug is required",
		"slug":     "Slug can only contain lowercase letters, numbers, and hyphens",
	},
	"ID": {
		"gt": "ID must be greater than 0",
	},
	"Search": {
		"required":      "Search is required",
		"alphanumspace": "Search can only contain letters, numbers, and spaces",
	},
	"UUID": {
		"uuid4": "Invalid UUID format bro",
	},
}
var formattedMessages = map[string]map[string]string{
	"Price": {
		"lte": "Price must be less than or equal to %s",
	},
	"Slug": {
		"min": "%s must be at least %s characters",
		"max": "%s must be at most %s characters",
	},
	"Search": {
		"min": "Search must be at least %s characters",
		"max": "Search must be at most %s characters",
	},
	"Lang": {
		"oneof": "Language must be one of: %s",
	},
}

// utils/validator.go
func FormatValidationErrors(err error) map[string]string {
	formatted := map[string]string{}

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			// Namespace: CreateProductRequest.Image[2].URL
			// StructNamespace: CreateProductRequest.Image[2]
			// We want: Image[2].URL
			ns := fe.Namespace()
			root := strings.Split(ns, ".")
			if len(root) > 1 {
				formatted[strings.Join(root[1:], ".")] = GetCustomErrorMessage(fe.Field(), fe.Tag(), fe)
			} else {
				formatted[fe.Field()] = GetCustomErrorMessage(fe.Field(), fe.Tag(), fe)
			}
		}
	}

	return formatted
}

func GetCustomErrorMessage(field, tag string, fe validator.FieldError) string {
	log.Printf("[Validation] Field: %s | Tag: %s | Param: %s", field, tag, fe.Param())
	if msg, ok := staticMessages[field][tag]; ok {
		log.Printf("[Validation] Message returned: %s", msg)
		return msg
	}

	if format, ok := formattedMessages[field][tag]; ok {
		switch tag {
		case "min", "max":
			if strings.Contains(format, "%s must") {
				return fmt.Sprintf(format, field, fe.Param())
			}
			return fmt.Sprintf(format, fe.Param())
		case "oneof":
			return fmt.Sprintf(format, fe.Param())
		}
	}

	return fmt.Sprintf("Invalid value for %s", field)
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
