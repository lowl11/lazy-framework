package exceptions

var (
	RouteNotFound = New(
		"Route not found",
		"Путь не найден",
	)

	Timeout = New(
		"Request reached timed out",
		"Время работы истекло",
	)
)
