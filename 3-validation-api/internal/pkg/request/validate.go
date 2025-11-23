package request

import "github.com/go-playground/validator/v10"

func IsValid[T any](payLoad T) error {
	validate := validator.New()
	err := validate.Struct(payLoad)
	return err
}
