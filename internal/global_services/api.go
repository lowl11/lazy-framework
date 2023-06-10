package global_services

import (
	"github.com/lowl11/lazy-framework/services/script_service"
	"github.com/lowl11/lazylog/log"
)

var (
	Script *script_service.Service
)

func InitScript() {
	var err error

	Script, err = script_service.New()
	if err != nil {
		log.Fatal(err, "Initialize Script service error")
	}
}
