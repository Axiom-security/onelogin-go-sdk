package models

import (
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
		return v != nil && *v != ""
	case string:
		return v != ""
	default:
		return false
	}
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
