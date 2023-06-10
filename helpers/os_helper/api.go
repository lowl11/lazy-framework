package os_helper

import (
	"github.com/lowl11/lazy-framework/log"
	"os"
	"strings"
)

func NoProxy(hosts ...string) {
	if len(hosts) == 0 {
		return
	}

	SetEnv("no_proxy", strings.Join(hosts, ","))
}

func Get(key string) string {
	return os.Getenv(key)
}

func SetEnv(key, value string) {
	lowercase := strings.ToLower(key)
	uppercase := strings.ToUpper(key)

	if err := os.Setenv(lowercase, value); err != nil {
		log.Error(err, "Set "+key+" env error")
	}

	if err := os.Setenv(uppercase, value); err != nil {
		log.Error(err, "Set "+key+" env error")
	}
}
