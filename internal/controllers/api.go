package controllers

import (
	"github.com/lowl11/owl/internal/controllers/static_controller"
)

var (
	Static *static_controller.Controller
)

func Init() {
	Static = static_controller.New()
}
