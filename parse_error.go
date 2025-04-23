package validify

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func (v *ValidatorImpl) ParseErrors(err interface{}) []string {
	exception, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	errorMessages := make([]string, 0)

	// Define error message templates for tags
	tagMessages := map[string]func(err validator.FieldError) string{
		"required": func(err validator.FieldError) string {
			return fmt.Sprintf("%s is required", err.Field())
		},
		"min": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
		},
		"max": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must be at most %s characters long", err.Field(), err.Param())
		},
		"eq": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must be equal to %s", err.Field(), err.Param())
		},
		"ne": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must not equal to %s", err.Field(), err.Param())
		},
		"gt": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must be greater than %s", err.Field(), err.Param())
		},
		"gte": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must be greater than or equal %s", err.Field(), err.Param())
		},
		"lt": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must be less than %s", err.Field(), err.Param())
		},
		"lte": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must be less than or equal %s", err.Field(), err.Param())
		},
		"isValidEmail": func(err validator.FieldError) string {
			return fmt.Sprintf("%s address is invalid", err.Field())
		},
		"isValidPhoneNumber": func(err validator.FieldError) string {
			return fmt.Sprintf("%s must start with 62xx", err.Field())
		},
	}

	// Add error message template for provided custom validations
	for name, item := range v.customValidation {
		tagMessages[name] = item.ErrorMessage
	}

	// Add error messages based on tags
	for _, e := range exception {
		if handler, exists := tagMessages[e.Tag()]; exists {
			// Handle predefined tags
			errorMessages = append(errorMessages, handler(e))
		} else {
			// Default case for unhandled tags
			errorMessages = append(errorMessages, fmt.Sprintf("%s is invalid", e.Field()))
		}
	}

	return errorMessages
}
