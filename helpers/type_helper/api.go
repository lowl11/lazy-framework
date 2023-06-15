package type_helper

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func ToString(anyValue any) string {
	if _, ok := anyValue.(error); ok {
		return anyValue.(error).Error()
	}

	value := reflect.ValueOf(anyValue)

	switch value.Kind() {
	case reflect.String:
		return anyValue.(string)
	case reflect.Bool:
		return strconv.FormatBool(anyValue.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32:
		return fmt.Sprintf("%f", value.Float())
	case reflect.Float64:
		return fmt.Sprintf("%g", value.Float())
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		valueInBytes, err := json.Marshal(anyValue)
		if err != nil {
			return ""
		}
		return string(valueInBytes)
	case reflect.Ptr:
		return ToString(value.Elem().Interface())
	default:
		return fmt.Sprintf("%v", value)
	}
}

func IsEmptySlice(value any) bool {
	return reflect.ValueOf(value).Len() == 0
}

func IsSlice(value any) bool {
	kind := reflect.TypeOf(value).Kind()
	return kind == reflect.Slice ||
		kind == reflect.Array
}
