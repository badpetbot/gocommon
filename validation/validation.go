package validation

import(

  // Import builtin packages.
  "regexp"

  // Import 3rd party packages.
  validator "github.com/go-playground/validator/v10"
)

// NewValidator creates and returns a new go-playground/validator.Validate with all custom validation functions, pre-registered.
func NewValidator() *validator.Validate {

  // Create a new validator.
  v := validator.New()

  // Register custom validations. This is an example copied from something else and probably won't actually be used.
  v.RegisterValidation("display_name", validateDisplayNameFactory())

  // Return it.
  return v
}

// ------------------------------- //
// ----- Validator Factories ----- //
// ------------------------------- //

// Validate a display name.
var displayNameRegex *regexp.Regexp
func validateDisplayNameFactory() validator.Func {

  displayNameRegex = regexp.MustCompile("^[a-zA-Z0-9-_]{5,30}$")

  return func(f validator.FieldLevel) bool {
    return displayNameRegex.MatchString(f.Field().String())
  }
}