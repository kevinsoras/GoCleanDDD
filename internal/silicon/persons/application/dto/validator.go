// internal/silicon/persons/application/dto/validator.go
package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/kevinsoras/GoCleanDDD/internal/silicon/shared/interfaces"
)

type customValidator struct {
	v *validator.Validate
}

func NewCustomValidator() interfaces.Validator {
	v := validator.New()
	cv := &customValidator{v: v}
	cv.v.RegisterStructValidation(CreatePersonInputStructLevelValidation, CreatePersonInput{})
	return cv
}

func (cv *customValidator) ValidateStruct(s interface{}) error {
	return cv.v.Struct(s)
}
