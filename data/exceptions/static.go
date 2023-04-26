package exceptions

import "net/http"

var (
	RouteNotFound = New(
		"Route not found",
		http.StatusNotFound,
	)

	Timeout = New(
		"Request reached timed out",
	)
)
