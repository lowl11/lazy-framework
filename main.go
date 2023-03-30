package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lowl11/lazy-framework/framework"
	"github.com/lowl11/lazy-framework/log"
	"net/http"
	"time"
)

func main() {
	//framework.SetLogConfig("info", "logs")
	framework.SetServerTimeout(time.Second * 30)

	if err := framework.Init(); err != nil {
		log.Fatal(err, "Framework init error")
	}

	setRoutes(framework.ServerEcho())

	framework.Server().Start(":8080")

	log.Info("hello world")
}

func setRoutes(server *echo.Echo) {
	server.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
