package validation

import (
	"errors"
)

func RequiredField(value any, name string) error {
	if value == nil {
		return errors.New("Field " + name + " is null, but it's required")
	}

	_, isString := value.(string)
	if isString && value.(string) == "" {
		return errors.New("Field " + name + " is null or empty, but it's required")
	}

	_, isInt := value.(int)
	if isInt && value.(int) == 0 {
		return errors.New("Field " + name + " is null or zero, but it's required")
	}

	return nil
}
