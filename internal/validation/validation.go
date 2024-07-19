package validation

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func GetValidationErrors(obj any, err error) map[string]string {
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
	return errs
}

func getMsg(params ...string) map[string]string {
	return map[string]string{
		"required": fmt.Sprintf("%s is required", params[0]),
		"email":    "Invalid email",
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
