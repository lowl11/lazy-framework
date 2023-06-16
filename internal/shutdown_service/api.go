package shutdown_service

func (service *Service) Add(action func()) *Service {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	service.actions = append(service.actions, action)
	return service
}

func (service *Service) Run() {
	for _, action := range service.actions {
		service.runFunc(action)
	}
}
