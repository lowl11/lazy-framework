package validation

import (
	"errors"
	"github.com/lowl11/lazy-framework/helpers/type_helper"
	"unicode"
)

func Required(value any, name string) error {
	if value == nil {
		return errors.New("Field " + name + " is null, but it's required")
	}

	_, isStringType := value.(string)
	if isStringType && value.(string) == "" {
		return newError(name)
	}

	_, isInt := value.(int)
	if isInt && value.(int) == 0 {
		return newNumeric(name)
	}

	_, isFloat32 := value.(float32)
	if isFloat32 && value.(float32) == 0 {
		return newNumeric(name)
	}

	if type_helper.IsSlice(value) && type_helper.IsEmptySlice(value) {
		return newError(name)
	}

	return nil
}

func IsPrimitive(value any) bool {
	return isInteger(value) || isString(value) || isBool(value) || isFloat(value)
}

func IsLower(value string) bool {
	for _, r := range value {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true

}

func IsUpper(value string) bool {
	for _, r := range value {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
