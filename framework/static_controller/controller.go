package static_controller

import "github.com/lowl11/lazy-framework/framework/controller"

type Controller struct {
	controller.Base
}

func Create() *Controller {
	return &Controller{}
}