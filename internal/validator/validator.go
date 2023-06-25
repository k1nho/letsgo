package validator

import "regexp"

var (
	emailRX = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\. [a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

// Define a new Validator type which contains a map of validation errors.
type Validator struct {
	Errors map[string]string
}

// New: return an instance of a validator
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid: returns true if there are no errors, otherwise false
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError: add a new error to the validator
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check: Given a bool expression, if it evaluates to false then the validator adds an error, otherwise it skips (validation is correct)
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// In: Returns true if a value exists in a list, otherwise false
func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}
	return false
}

func Matches(value string, rg *regexp.Regexp) bool {
	return rg.MatchString(value)
}

// Unique: returns true if all the values are unique, otherwise false
func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)

	for _, val := range values {
		uniqueValues[val] = true
	}
	return len(values) == len(uniqueValues)
}
