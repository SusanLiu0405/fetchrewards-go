package models

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required,alphanum"`
	Price            string `json:"price" validate:"required,price"`
}

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("price", validatePrice)
	return validate
}

func validatePrice(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^\d+(\.\d{2})?$`)
	return re.MatchString(fl.Field().String())
}
