package grpc_service

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

func (service *Service) Credentials(creds credentials.TransportCredentials) *Service {
	service.creds = creds
	return service
}

func (service *Service) Options(dialOptions ...grpc.DialOption) *Service {
	service.opts = append(service.opts, dialOptions...)
	return service
}

func (service *Service) Timeout(duration time.Duration) *Service {
	service.timeout = duration
	return service
}

func (service *Service) NoProxy() *Service {
	service.noProxy = true
	return service
}

func (service *Service) SslTrust() *Service {
	service.sslCheck = true
	return service
}

func (service *Service) Connection() (*grpc.ClientConn, error) {
	connection, err := service.connect()
	if err != nil {
		return nil, err
	}

	return connection, nil
}
