package validation

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type Validation struct {
	validator *validator.Validate
	Details   any
}

func NewValidator() *Validation {
	validator := validator.New()

	return &Validation{
		validator: validator,
	}
}

func (v *Validation) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return getValidationErrors(i, err)
	}
	return nil
}

func (ves Validation) Error() string {
	return "validation errors"
}

func getValidationErrors(obj any, err error) error {
	errs := make(map[string]string)
	var ves validator.ValidationErrors
	if errors.As(err, &ves) {
		for _, err := range ves {
			fieldName := getName(obj, err.Field())
			if msg, ok := getMsg(fieldName, err.Param())[err.Tag()]; ok {
				errs[fieldName] = msg
			}
		}
	}
	if len(errs) == 0 {
		return err
	}
	return Validation{Details: errs}
}

func getMsg(params ...string) map[string]string {
	return map[string]string{
		"required": fmt.Sprintf("%s is required", params[0]),
		"email":    "Invalid email",
		"gt":    fmt.Sprintf("%s is greater than %s", params[0], params[1]),
	}
}

func getName(obj any, fieldName string) string {
	field, ok := reflect.TypeOf(obj).FieldByName(fieldName)
	if ok {
		tag := field.Tag.Get("json")
		if tag != "" && tag != "-" {
			return tag
		}
	}
	return fieldName
}
