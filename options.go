package validify

// Option -.
type Option func(*ValidatorImpl)

// WithCustomValidation -.
func WithCustomValidation(customValidation map[string]CustomValidation) Option {
	return func(c *ValidatorImpl) {
		c.customValidation = customValidation
	}
}
