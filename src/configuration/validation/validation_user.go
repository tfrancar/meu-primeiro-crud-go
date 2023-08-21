package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/tfrancar/meu-primeiro-crud-go/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	tradutor ut.Translator
)

func init() {
	//no momento que estou usando o .(*validator.validate) estou fazendo um casting
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		tradutor, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, tradutor)

	}

}

func ValidateUserError(validation_err error) *rest_err.RestErr {

	var (
		jsonErr             *json.UnmarshalTypeError
		jsonValidationError validator.ValidationErrors
	)

	// Caso o erro seja um tipo de erro que não precise converter
	// ex.: mandou um string no lugar de um número ou vice-versa
	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type") //Campo incorreto.
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{} //caso seja mais de campo errado, será preenchido nesta variável.

		for _, e := range validation_err.(validator.ValidationErrors) {
			causes := rest_err.Causes{
				Message: e.Translate(tradutor),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, causes)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)

	} else {

		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
