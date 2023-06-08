package grpc_service

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func (service *Service) connect() (*grpc.ClientConn, error) {
	//creds := insecure.NewCredentials()
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	if service.creds != nil {
		creds = service.creds
	}

	options := service.opts
	options = append(options, grpc.WithTransportCredentials(creds))

	// set no proxy
	if service.noProxy {
		options = append(options, grpc.WithNoProxy())
	}

	ctx, cancel := context.WithTimeout(context.Background(), service.timeout)
	defer cancel()

	connection, err := grpc.DialContext(ctx, service.host, options...)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

func (service *Service) asd() {
	//
}
