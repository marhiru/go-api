package validation

import (
	"api/src/configuration/rest_err"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate   = validator.New()
	translator ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		translator, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, translator)
	}
}

func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(translator),
				Field:   e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
