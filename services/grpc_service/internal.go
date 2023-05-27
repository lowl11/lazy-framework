package grpc_service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (service *Service) connect() (*grpc.ClientConn, error) {
	creds := insecure.NewCredentials()
	if service.creds != nil {
		creds = service.creds
	}

	options := service.opts
	options = append(options, grpc.WithTransportCredentials(creds))

	ctx, cancel := context.WithTimeout(context.Background(), service.timeout)
	defer cancel()

	connection, err := grpc.DialContext(ctx, service.host, options...)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
