package shutdown_service

import "sync"

type Service struct {
	mutex   sync.Mutex
	actions []func()
}

func New() *Service {
	return &Service{
		actions: make([]func(), 0, 3),
	}
}
