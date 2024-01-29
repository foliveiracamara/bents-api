package validation

import (
	"encoding/json"
	"errors"

	"github.com/foliveiracamara/bents-api/configuration/apperr"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	transl ut.Translator
)

func init() {
	val := validator.New()
	enTranslator := en.New()
	unt := ut.New(enTranslator, enTranslator)

	var err error
	transl, ok := unt.GetTranslator("en")
	if !ok {
		panic("Failed to get translator")
	}

	err = en_translation.RegisterDefaultTranslations(val, transl)
	if err != nil {
		panic(err)
	}
}

func ValidateUserError(validation_err error) *apperr.AppErr {
	var jsonErr *json.UnsupportedTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return apperr.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []apperr.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := apperr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return apperr.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return apperr.NewBadRequestError("Error trying to convert fields")
	}
}
