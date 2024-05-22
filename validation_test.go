package belajar_golang_validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validation instance is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()
	var nilString string
	err := validate.Var(nilString, "required")
	if err != nil {
		t.Error(err)
	}
}

func TestValidateTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "rahasia"

	if err := validate.VarWithValue(password, confirmPassword, "eqfield"); err != nil {
		t.Error(err)
	}
}

// for more tags in validator, go check baked-in validation on the doc of validator package

func TestMultipleTagValidation(t *testing.T) {
	validate := validator.New()
	user := "12345"

	if err := validate.Var(user, "required,numeric"); err != nil {
		t.Error(err)
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()
	user := "1111111111"

	if err := validate.Var(user, "required,numeric,min=5,max=10"); err != nil {
		t.Error(err)
	}
}

func TestStructValidation(t *testing.T) {
	validate := validator.New()
	user := struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}{
		Username: "dendun",
		Password: "123",
	}

	if err := validate.Struct(user); err != nil {
		//t.Error(err.Error())
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		for _, fieldError := range validationErrors {
			//fmt.Println(fieldError.Field(), fieldError.Tag(), fieldError.Value())
			t.Error("Error:", fieldError.Field(), "tag_violation:", fieldError.Tag(), "value:", fieldError.Value())
		}
		//t.Error(validationErrors)
	}
}
