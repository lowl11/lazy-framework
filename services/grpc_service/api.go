package grpc_service

import (
	"github.com/lowl11/lazy-framework/framework"
	"github.com/lowl11/lazy-framework/log"
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

func (service *Service) Singleton() (*grpc.ClientConn, error) {
	connection, err := service.connect()
	if err != nil {
		return nil, err
	}

	framework.ShutDownAction(func() {
		if err = connection.Close(); err != nil {
			log.Error(err, "Close gRPC client connection error")
			return
		}
		log.Info("gRPC client connection closed!")
	})

	return connection, nil
}

func (service *Service) Connection() (*grpc.ClientConn, error) {
	connection, err := service.connect()
	if err != nil {
		return nil, err
	}

	return connection, nil
}
