package models

import (
	"strconv"
	"strings"
	"time"
)

type Queryable interface {
	GetKeyValidators() map[string]func(interface{}) bool
	SetLimit(limit string)
	SetPage(page string)
	SetCursor(cursor string)
}

// validateString checks if the value is a valid string
func validateString(value interface{}) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case *string:
		return v != nil && len(strings.TrimSpace(*v)) > 0
	case string:
		return len(strings.TrimSpace(v)) > 0
	default:
		return false
	}
}

// validateCommaSeparatedList checks if the value is a valid comma-separated list
// where each item is non-empty after trimming whitespace
func validateCommaSeparatedList(value interface{}) bool {
	if !validateString(value) {
		return false
	}

	var str string
	switch v := value.(type) {
	case *string:
		str = *v
	case string:
		str = v
	default:
		return false
	}

	// Split by comma and check each item
	items := strings.Split(str, ",")
	for _, item := range items {
		if len(strings.TrimSpace(item)) == 0 {
			return false
		}
	}
	return true
}

// validateNumericString checks if the value is a valid numeric string
func validateNumericString(value interface{}) bool {
	if !validateString(value) {
		return false
	}

	var str string
	switch v := value.(type) {
	case *string:
		str = *v
	case string:
		str = v
	default:
		return false
	}

	// Check if it's a valid integer
	str = strings.TrimSpace(str)
	if _, err := strconv.ParseInt(str, 10, 64); err != nil {
		return false
	}
	return true
}

// validateTime checks if the value is a valid time.Time
func validateTime(value interface{}) bool {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case *time.Time:
		return v != nil && !v.IsZero()
	case time.Time:
		return !v.IsZero()
	default:
		return false
	}
}

// validateInt checks if the provided value is an int.
func validateInt(val interface{}) bool {
	switch v := val.(type) {
	case int:
		return true
	case *int:
		return v != nil
	default:
		return false
	}
}

// validateBool checks if the provided value is a bool.
func validateBool(val interface{}) bool {
	switch v := val.(type) {
	case bool:
		return true
	case *bool:
		return v != nil
	default:
		return false
	}
}
