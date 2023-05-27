package grpc_service

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Service struct {
	host  string
	creds credentials.TransportCredentials
	opts  []grpc.DialOption
}

func New(host string) *Service {
	return &Service{
		host: host,
		opts: make([]grpc.DialOption, 0),
	}
}
