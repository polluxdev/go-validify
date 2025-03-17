package govalidator

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Functions() map[string]validator.Func
	RegisterCustomValidator()
	Validate(req interface{}) error
	ParseErrors(err interface{}) []string
}

type ValidatorImpl struct {
	validate *validator.Validate
}

func New() Validator {
	return &ValidatorImpl{
		validate: validator.New(),
	}
}

func (v *ValidatorImpl) Functions() map[string]validator.Func {
	return map[string]validator.Func{
		"isValidEmail":       isValidEmail,
		"isValidPhoneNumber": isValidPhoneNumber,
	}
}

func (v *ValidatorImpl) RegisterCustomValidator() {
	for name, fn := range v.Functions() {
		err := v.validate.RegisterValidation(name, fn)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (v *ValidatorImpl) Validate(req interface{}) error {
	return v.validate.Struct(req)
}
