# lazy-framework

> Collection of other "lowl11" libraries need to build web microservices

## Example

Start simple REST server with Echo library
```go
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/framework"
	"github.com/lowl11/lazy-framework/log"
	"net/http"
	"time"
)

func main() {
	// custom settings before calling initialization (in .Server().Start())
	framework.SetLogConfig("info", "logs")       // setting custom logger config
	framework.SetServerTimeout(time.Second * 30) // setting custom server timeout (REST)

	// setting Echo routes
	setRoutes(framework.ServerEcho())

	// global logging package
	log.Info("hello world")

	// starting server (calls .Fatal() log if catch error)
	framework.Server().Start(":8080")
}

func setRoutes(server *echo.Echo) {
	server.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
```