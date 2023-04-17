package filter

import "strings"

func String(value string) string {
	return strings.TrimSpace(strings.ToLower(value))
}

func StringSimple(value string) string {
	return strings.TrimSpace(value)
}
