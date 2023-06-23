# lazy

> Collection of other "lowl11" libraries need to build web microservices

### Example
<hr>

Start simple REST server with Echo library
```go
package main

import (
	"github.com/lowl11/owl"
	"github.com/lowl11/lazyconfig/config"
)

func main() {
	// initialize server
	server := owl.New(&owl.Config{})

	// run the server
	server.Start(config.Get("server_port"))
}
```

### Logging
<hr>

Call log from anywhere!
```go
package main

import "github.com/lowl11/owllog/log"

func main() {
	log.Debug("debug message")
	log.Info("info message", 1, true, false)
	log.Warn("warning message")
	log.Error(err, "some error message")
	log.Fatal(err, "some fatal message")
}
```

Set custom logger with common [interface](https://github.com/lowl11/lazylog/blob/master/logapi/interface.go)

Example:
```go
package main

import "github.com/lowl11/owl"

func main() {
    // initialize server 
    server := owl.New(&owl.Config{
        CustomLoggers: []logapi.ILogger{
            &myLogger,
        },
    })

    // run the server
    server.Start(config.Get("server_port"))
}
```

### Config
<hr>

Reads config values from profiles/<env_name>.yml file. <br>
If yml file has no value for key, it goes from ENV
```go
package main

import (
	"fmt"
	"github.com/lowl11/owlconfig/config"
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
:white_check_mark: REST <br>
:white_check_mark: Cron Job <br>
:white_check_mark: SOAP <br>
:white_check_mark: gRPC <br>