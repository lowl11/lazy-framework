package filter

import "strings"

func String(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func StringSimple(value string) string {
	return strings.TrimSpace(value)
}

func PtrValue(value *string) string {
	if value != nil {
		return *value
	}

	return ""
}
