package main

import (
	"github.com/lowl11/lazy-framework/framework"
	"github.com/lowl11/lazy-framework/log"
)

func main() {
	framework.SetLogConfig("info", "logs")

	if err := framework.Init(); err != nil {
		log.Fatal(err, "Framework init error: ")
	}

	log.Info("hello world")
}
