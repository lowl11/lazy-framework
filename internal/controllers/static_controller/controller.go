package static_controller

import (
	"github.com/lowl11/lazy-framework/internal/base/controller"
)

type Controller struct {
	controller.Base
}

func New() *Controller {
	return &Controller{}
}
