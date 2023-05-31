package validation

import (
	"errors"
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

	return nil
}

func IsPrimitive(value any) bool {
	return isInteger(value) || isString(value) || isBool(value) || isFloat(value)
}
