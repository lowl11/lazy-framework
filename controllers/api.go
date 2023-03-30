package controllers

import "github.com/lowl11/lazy-framework/controllers/static_controller"

var (
	Static *static_controller.Controller
)

func Init() {
	Static = static_controller.Create()
}
