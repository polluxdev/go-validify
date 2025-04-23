package main

import (
	"fmt"
	"reflect"
	"strings"

	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/polluxdev/go-validify"
)

type Request struct {
	ID    string `validate:"required"`
	Name  string `validate:"required,min=5"`
	Email string `validate:"isValidEmail"`
	Type  string `validate:"isValidType"`
}

func main() {
	types := []string{"active", "inactive"}

	customFunctions := map[string]validify.CustomValidation{
		"isValidType": {
			Function: func(fl validator.FieldLevel) bool {
				if fl.Field().Kind() != reflect.String {
					return false
				}

				field := fl.Field().String()

				return slices.Contains(types, field)
			},
			ErrorMessage: func(err validator.FieldError) string {
				return fmt.Sprintf("%s must be one of %s", err.Field(), strings.Join(types, ", "))
			},
		},
	}

	validator := validify.New(validify.WithCustomValidation(customFunctions))
	validator.RegisterCustomValidator()

	request := Request{
		Name:  "bob",
		Email: "invalid",
	}

	err := validator.Validate(request)
	if err != nil {
		errMsg := validator.ParseErrors(err)
		for _, msg := range errMsg {
			fmt.Println(msg)
		}
	}
}
