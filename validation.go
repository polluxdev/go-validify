package validify

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func isValidEmail(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}

	email := fl.Field().String()
	pattern := "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$"
	return regexp.MustCompile(pattern).MatchString(email)
}

func isValidPhoneNumber(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}

	field := fl.Field().String()
	pattern := "^62[1-9][1-9]"
	return regexp.MustCompile(pattern).MatchString(field)
}
