package validator

import "github.com/go-playground/validator/v10"

var valid *validator.Validate

// init will be run when the package is started,
// this way you avoid having to recreate an important variable every time the package is used.
func init() {
	valid = validator.New()
}

// ValidStruct performs validation of a data structure using the validator.
// Receives an arbitrary data structure and uses the validator to verify
// if it meets the criteria defined in the validation tags.
// Returns an error if validation fails, indicating fields that are not compliant.
func ValidStruct(data any) error {
	return valid.Struct(data)
}
