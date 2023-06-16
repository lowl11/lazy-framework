package shutdown_service

import (
	"errors"
	"github.com/lowl11/lazylog/log"
)

func (service *Service) runFunc(action func()) {
	defer func() {
		if value := recover(); value != nil {
			var err error

			if _, ok := value.(string); ok {
				err = errors.New(value.(string))
			} else if _, ok = value.(error); ok {
				err = value.(error)
			}

			log.Error(err, "Catch panic from shut down action")
		}
	}()

	action()
}
