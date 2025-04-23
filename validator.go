package validify

import (
	"log"

	"github.com/go-playground/validator/v10"
)

type Validator interface {
	RegisterCustomValidator()
	Validate(req any) error
	ParseErrors(err any) []string
}

type CustomValidation struct {
	Function     validator.Func
	ErrorMessage func(err validator.FieldError) string
}

type ValidatorImpl struct {
	validate         *validator.Validate
	customValidation map[string]CustomValidation
}

func New(opts ...Option) Validator {
	v := &ValidatorImpl{
		validate: validator.New(),
	}

	for _, opt := range opts {
		opt(v)
	}

	return v
}

func (v *ValidatorImpl) functions() map[string]validator.Func {
	return map[string]validator.Func{
		"isValidEmail":       isValidEmail,
		"isValidPhoneNumber": isValidPhoneNumber,
	}
}

func (v *ValidatorImpl) RegisterCustomValidator() {
	// Register built-in validations
	for name, fn := range v.functions() {
		err := v.validate.RegisterValidation(name, fn)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Register provided custom validations
	for name, item := range v.customValidation {
		if err := v.validate.RegisterValidation(name, item.Function); err != nil {
			log.Fatal(err)
		}
	}
}

func (v *ValidatorImpl) Validate(req any) error {
	return v.validate.Struct(req)
}
