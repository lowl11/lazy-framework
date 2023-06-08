package grpc_service

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

type Service struct {
	host     string
	creds    credentials.TransportCredentials
	opts     []grpc.DialOption
	timeout  time.Duration
	noProxy  bool
	sslCheck bool
}

func New(host string) *Service {
	return &Service{
		host:    host,
		opts:    make([]grpc.DialOption, 0),
		timeout: time.Second * 30,
	}
}
