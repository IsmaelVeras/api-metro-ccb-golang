package validation

import (
	"encoding/json"
	"errors"

	rest_err "github.com/IsmaelVeras/api-golang-crud/src/configuration"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_BR_translation "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	Validation = validator.New()
	Transl     ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		Transl, _ = unt.GetTranslator("en")
		pt_BR_translation.RegisterDefaultTranslations(val, Transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidatonError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Tipo de campo inválido")
	} else if errors.As(validation_err, &jsonValidatonError) {

		errorsCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(Transl),
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Alguns campos são inválidos", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Erro ao tentar validar usuário")
	}
}
