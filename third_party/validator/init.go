package validator

import (
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"sync"
)

var once sync.Once

type Util interface {
	Validate() *validator.Validate
	Struct(payload interface{}) error
}

type util struct {
	Validator *validator.Validate
}

func Init() Util {
	var validate *validator.Validate

	once.Do(func() {
		validate = validator.New()
	})

	return &util{
		Validator: validate,
	}
}

func (u util) Validate() *validator.Validate {
	return u.Validator
}

func (u util) Struct(payload interface{}) error {
	//validate valid tag
	if err := u.Validator.Struct(payload); err != nil {
		errorMessages := ""

		for _, err := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprint(err.Field(), " with value '", err.Value(), "' not valid for '", err.Tag(), err.Param(), "' validation. ")
			errorMessages += errorMessage
		}

		return errors.New(errorMessages)
	}
	return nil
}
