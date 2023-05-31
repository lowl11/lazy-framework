package validation

import (
	"errors"
	"reflect"
)

func newError(name string) error {
	return errors.New("Field " + name + " is null or empty, but it's required")
}

func newNumeric(name string) error {
	return errors.New("Field " + name + " is null or zero, but it's required")
}

func isInteger(value any) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
		return true
	}
	return false
}

func isFloat(value any) bool {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
	case reflect.Complex128:
		return true
	}
	return false
}

func isBool(value any) bool {
	_, ok := value.(bool)
	return ok
}

func isString(value any) bool {
	_, ok := value.(string)
	return ok
}
