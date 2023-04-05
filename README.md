# lazy-framework

> Collection of other "lowl11" libraries need to build web microservices

### Example
<hr>

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
	// turn on swagger endpoints
	framework.UseSwagger()

	// set config & environment settings
	framework.SetEnvironmentDefault("local")

	// turn on HTTP 2.0 mode
	framework.UseHttp2(&models.Http2Config{
		MaxConcurrentStreams: 123,
		MaxReadFrameSize:     123,
	})

	// setting custom loggers
	framework.SetCustomLoggers( /* some custom loggers */ )

	// custom settings before calling initialization (in .Server().Start())
	framework.SetLogConfig("info", "logs")       // setting custom logger config
	framework.SetServerTimeout(time.Second * 30) // setting custom server timeout (REST)

	// init framework with given settings
	framework.Init()

	// setting Echo routes
	setRoutes(framework.ServerEcho())

	// print value from yaml config
	log.Info("value from env:", config.Get("test_key"))

	// global logging package
	log.Info("hello world")

	log.Info("Environment:", config.Env())

	// starting server (calls .Fatal() log if catch error)
	framework.Server().Start(":8080")
}

func setRoutes(server *echo.Echo) {
	server.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "test endpoint")
	})
}
```

### Logging
<hr>

Call log from anywhere!
```go
package main

import "github.com/lowl11/lazy-framework/log"

func main() {
	log.Info("test message", 1, true, false)
}
```

Set custom logger with common [interface](https://github.com/lowl11/lazylog/blob/master/logapi/interface.go). For example, for ElasticSearch

Example:
```go
package main

import "github.com/lowl11/lazy-framework/framework"

func main() {
	framework.SetCustomLoggers(someCustomLogger1, someCustomLogger2)
}
```

### Config
<hr>

Reads config values from profiles/<env_name>.yml file.
```go
package main

import (
	"fmt"
	"github.com/lowl11/lazy-framework/config"
)

func main() {
	testKeyValue := config.Get("test_key")
	fmt.Println(testKeyValue)
}
```

Also replace ENV variables given in .yml file. Example:
```yaml
test_key: test value 123
test_env_key: {{TEST_ENV_KEY}} # <- will be replaced from env
```


### Features
<hr>

Web Frameworks <br>
:white_check_mark: Echo <br>
:x: Gin <br>